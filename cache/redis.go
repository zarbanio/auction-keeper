package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/entities"

	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	client redis.UniversalClient
}

// GetAddresses implements ICache.
func (*redisStore) GetAddresses(ctx context.Context) (map[string]common.Address, error) {
	panic("unimplemented")
}

// SaveAddresses implements ICache.
func (*redisStore) SaveAddresses(ctx context.Context, addresses map[string]common.Address) error {
	panic("unimplemented")
}

func NewRedisStore(addr string) ICache {
	return &redisStore{
		client: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (r redisStore) SaveVatIlk(ctx context.Context, ilk entities.VatIlk, expiration time.Duration) error {
	value, err := json.Marshal(ilk)
	if err != nil {
		return nil
	}
	return r.client.Set(ctx, fmt.Sprintf("ilks:%s", ilk.Name), value, expiration).Err()
}

func (r redisStore) GetVatIlkByName(ctx context.Context, ilkName string) (*entities.VatIlk, error) {
	res, err := r.client.Get(ctx, fmt.Sprintf("ilks:%s", ilkName)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrIlkNotFound
		}
		return nil, err
	}
	ilk := &entities.VatIlk{}
	err = json.Unmarshal([]byte(res), &ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting ilks. %w", err)
	}
	return ilk, nil
}

func (r redisStore) SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error {
	value, err := json.Marshal(vault)
	if err != nil {
		return nil
	}
	return r.client.Set(ctx, fmt.Sprintf("vaults:ilk:%s:urn:%s", vault.IlkName, vault.UrnAddress), value, expiration).Err()
}

func (r redisStore) GetVaults(ctx context.Context) ([]*entities.Vault, error) {
	keys, err := r.client.Keys(ctx, "vaults:ilk:*:urn:*").Result()
	if err != nil {
		return nil, err
	}

	var vaults []*entities.Vault
	for _, key := range keys {
		vault, err := r.getVaultByKey(ctx, key)
		if err != nil {
			return nil, err
		}
		vaults = append(vaults, vault)
	}
	return vaults, nil
}

func (r redisStore) getVaultByKey(ctx context.Context, key string) (*entities.Vault, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrVaultNotFound
		}
		return nil, err
	}
	v := &entities.Vault{}
	err = json.Unmarshal([]byte(res), &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r redisStore) SaveEra(ctx context.Context, era *big.Int, expiration time.Duration) error {
	value, err := json.Marshal(era)
	if err != nil {
		return nil
	}
	return r.client.Set(ctx, fmt.Sprintf("Era:%s", era.String()), value, expiration).Err()
}

func (r redisStore) DeleteEra(ctx context.Context, era *big.Int) error {
	return r.client.Del(ctx, fmt.Sprintf("Era:%s", era.String())).Err()
}

func (r redisStore) GetEras(ctx context.Context) ([]*big.Int, error) {
	keys, err := r.client.Keys(ctx, "Era:*").Result()
	if err != nil {
		return nil, err
	}

	eras := make([]*big.Int, len(keys))
	for _, key := range keys {
		res, err := r.client.Get(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return nil, ErrEraNotFound
			}
			return nil, err
		}
		var era big.Int
		err = json.Unmarshal([]byte(res), &era)
		if err != nil {
			return nil, err
		}
		eras = append(eras, &era)
	}
	return eras, nil
}
