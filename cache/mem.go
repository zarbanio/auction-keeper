package cache

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type memcache struct {
	mem map[string]interface{}
	mu  *sync.Mutex
}

func NewMemCache() MemCache {
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
