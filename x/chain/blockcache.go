package chain

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockCache struct {
	client   *ethclient.Client
	cache    map[uint64]*types.Block
	cacheMux sync.Mutex
}

func NewBlockCache(eth *ethclient.Client) *BlockCache {
	return &BlockCache{
		client: eth,
		cache:  make(map[uint64]*types.Block),
	}
}

func (bc *BlockCache) GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	bc.cacheMux.Lock()
	defer bc.cacheMux.Unlock()

	block, ok := bc.cache[number]
	if ok {
		return block, nil
	}

	block, err := bc.client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return nil, err
	}

	bc.cache[number] = block
	return block, nil
}
