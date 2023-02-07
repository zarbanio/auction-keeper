package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	client redis.UniversalClient
}

func NewRedisStore(addr string) ICache {
	return &redisStore{
		client: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (r redisStore) SaveIlk(ctx context.Context, ilk entities.Ilk) error {
	value, err := json.Marshal(ilk)
	if err != nil {
		return nil
	}
	return r.client.Set(ctx, fmt.Sprintf("ilks:%s", ilk.Name), value, 0).Err()
}

func (r redisStore) GetIlkById(ctx context.Context, id [32]byte) (*entities.Ilk, error) {
	res, err := r.client.Get(ctx, fmt.Sprintf("ilks:%s", id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrIlkNotFound
		}
		return nil, err
	}
	v := &entities.Ilk{}
	err = json.Unmarshal([]byte(res), &v)
	if err != nil {
		return nil, fmt.Errorf("error getting ilks. %w", err)
	}
	return v, nil
}

func (r redisStore) SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error {
	value, err := json.Marshal(vault)
	if err != nil {
		return nil
	}
	return r.client.Set(ctx, fmt.Sprintf("vaults:ilk:%s:urn:%s", vault.Ilk.Name, vault.UrnAddress), value, expiration).Err()
}

func (r redisStore) GetVaultByIlkUrn(ctx context.Context, ilk [32]byte, urn common.Address) (*entities.Vault, error) {
	return r.getVaultByKey(ctx, fmt.Sprintf("vaults:ilk:%s:urn:%s", ilk, urn))
}

func (r redisStore) GetVaults(ctx context.Context) ([]entities.Vault, error) {
	keys, err := r.client.Keys(ctx, "vaults:ilk:*:urn:*").Result()
	if err != nil {
		return nil, err
	}

	var vaults []entities.Vault
	for _, key := range keys {
		vault, err := r.getVaultByKey(ctx, key)
		if err != nil {
			return nil, err
		}
		vaults = append(vaults, *vault)
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
