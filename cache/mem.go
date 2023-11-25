package cache

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type memcache struct {
	mem map[string]interface{}
	mu  *sync.Mutex
}

// DeleteEra implements ICache.
func (*memcache) DeleteEra(ctx context.Context, era *big.Int) error {
	panic("unimplemented")
}

// GetEras implements ICache.
func (*memcache) GetEras(ctx context.Context) ([]*big.Int, error) {
	panic("unimplemented")
}

// GetVatIlkByName implements ICache.
func (*memcache) GetVatIlkByName(ctx context.Context, ilkName string) (*entities.VatIlk, error) {
	panic("unimplemented")
}

// GetVaults implements ICache.
func (*memcache) GetVaults(ctx context.Context) ([]*entities.Vault, error) {
	panic("unimplemented")
}

// SaveEra implements ICache.
func (*memcache) SaveEra(ctx context.Context, era *big.Int, expiration time.Duration) error {
	panic("unimplemented")
}

// SaveVatIlk implements ICache.
func (*memcache) SaveVatIlk(ctx context.Context, ilk entities.VatIlk, expiration time.Duration) error {
	panic("unimplemented")
}

// SaveVault implements ICache.
func (*memcache) SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error {
	panic("unimplemented")
}

// getVaultByKey implements ICache.
func (*memcache) getVaultByKey(ctx context.Context, key string) (*entities.Vault, error) {
	panic("unimplemented")
}

func NewMemCache() ICache {
	return &memcache{
		mem: make(map[string]interface{}),
		mu:  &sync.Mutex{},
	}
}

func (m *memcache) SaveAddresses(ctx context.Context, addresses map[string]common.Address) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mem["addresses"] = addresses
	return nil
}

func (m *memcache) GetAddresses(ctx context.Context) (map[string]common.Address, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.mem["addresses"]
	if !ok {
		return nil, ErrAddressedNotFound
	}
	res := v.(map[string]common.Address)
	return res, nil
}
