package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain"
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

func (m *memcache) SaveProtocolData(ctx context.Context, protocolData domain.ProtocolData, expirationTime time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mem["protocol:data"] = protocolData
	return nil
}

func (m *memcache) GetProtocolData(ctx context.Context) (*domain.ProtocolData, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.mem["protocol:data"]
	if !ok {
		return nil, ErrProtocolDataNotFound
	}
	res := v.(domain.ProtocolData)
	return &res, nil
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

func (m *memcache) SaveProxy(ctx context.Context, proxy domain.Proxy) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mem[fmt.Sprintf("proxies:%s", proxy.Owner)] = proxy
	return nil
}

func (m *memcache) GetProxyByOwner(ctx context.Context, owner common.Address) (*domain.Proxy, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.mem[fmt.Sprintf("proxies:%s", owner)]
	if !ok {
		return nil, ErrProxyNotFound
	}
	res := v.(domain.Proxy)
	return &res, nil
}
