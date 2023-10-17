package cachedeth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type ethProxy struct {
	eth    *ethclient.Client
	bcache *BlockCache
	s      store.IStore
}

func NewEthProxy(eth *ethclient.Client, s store.IStore, bcache *BlockCache) eth.CachedEthereum {
	return &ethProxy{eth: eth, s: s, bcache: bcache}
}

func (e *ethProxy) CachedFilterLogs(context context.Context, q ethereum.FilterQuery) ([]eth.Log, error) {
	if len(q.Topics) > 0 {
		return nil, errors.New("topics are not supported")
	}
	if q.BlockHash != nil {
		return nil, errors.New("block hash is not supported")
	}

	lq := store.EthereumLogsQuery{
		FromBlock: q.FromBlock.Uint64(),
		ToBlock:   q.ToBlock.Uint64(),
		Addresses: q.Addresses,
	}

	isCached, err := e.s.IsEthereumLogsCached(context, lq)
	if err != nil {
		return nil, fmt.Errorf("error checking logs cached. %w", err)
	}
	if !isCached {
		logs, err := e.eth.FilterLogs(context, q)
		if err != nil {
			return nil, err
		}
		llogs := make([]eth.Log, 0, len(logs))
		for _, l := range logs {
			block, err := e.bcache.GetBlockByNumber(context, l.BlockNumber)
			if err != nil {
				return nil, err
			}
			log := eth.Log{
				Log:       l,
				Timestamp: time.Unix(int64(block.Header().Time), 0),
			}
			id, err := e.s.CreateLog(context, log)
			if err != nil {
				return nil, err
			}
			log.Id = id
			llogs = append(llogs, log)
		}

		err = e.s.CreateEthereumLogsCacheKey(context, lq)
		if err != nil {
			return nil, fmt.Errorf("error create logs cache key. %w", err)
		}
		return llogs, nil
	}

	logs, err := e.s.GetLogsByQuery(context, lq)
	if err != nil {
		return nil, fmt.Errorf("error getting logs with query. %w", err)
	}
	return logs, nil
}

func (e *ethProxy) CachedSubscribeFilterLogs(context context.Context, q ethereum.FilterQuery, ch chan<- eth.Log) (ethereum.Subscription, error) {
	logsCh := make(chan types.Log)
	sub, err := e.eth.SubscribeFilterLogs(context, q, logsCh)

	go func() {
		for l := range logsCh {
			block, err := e.bcache.GetBlockByNumber(context, l.BlockNumber)
			if err != nil {
				sub.Unsubscribe()
				return
			}
			l := eth.Log{
				Log:       l,
				Timestamp: time.Unix(int64(block.Header().Time), 0),
			}
			id, err := e.s.CreateLog(context, l)
			if err != nil {
				log.Printf("[CachedSubscribeFilterLogs] error creating log. %s.\n", err)
				sub.Unsubscribe()
				return
			}
			l.Id = id
			ch <- l
		}
	}()
	return sub, err
}

func (e *ethProxy) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return e.eth.BalanceAt(ctx, account, blockNumber)
}

func (e *ethProxy) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return e.eth.BlockByHash(ctx, hash)
}

func (e *ethProxy) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return e.eth.BlockByNumber(ctx, number)
}

func (e *ethProxy) BlockNumber(ctx context.Context) (uint64, error) {
	return e.eth.BlockNumber(ctx)
}

func (e *ethProxy) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return e.eth.CallContract(ctx, call, blockNumber)
}

func (e *ethProxy) ChainID(ctx context.Context) (*big.Int, error) {
	return e.eth.ChainID(ctx)
}

func (e *ethProxy) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return e.eth.CodeAt(ctx, account, blockNumber)
}

func (e *ethProxy) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return e.eth.EstimateGas(ctx, call)
}

func (e *ethProxy) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return e.eth.FilterLogs(ctx, q)
}

func (e *ethProxy) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return e.eth.HeaderByHash(ctx, hash)
}

func (e *ethProxy) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return e.eth.HeaderByNumber(ctx, number)
}

func (e *ethProxy) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return e.eth.NonceAt(ctx, account, blockNumber)
}

func (e *ethProxy) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	return e.eth.PendingBalanceAt(ctx, account)
}

func (e *ethProxy) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return e.eth.PendingCallContract(ctx, call)
}

func (e *ethProxy) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return e.eth.PendingCodeAt(ctx, account)
}

func (e *ethProxy) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return e.eth.PendingNonceAt(ctx, account)
}

func (e *ethProxy) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	return e.eth.PendingStorageAt(ctx, account, key)
}

func (e *ethProxy) PendingTransactionCount(ctx context.Context) (uint, error) {
	return e.eth.PendingTransactionCount(ctx)
}

func (e *ethProxy) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return e.eth.SendTransaction(ctx, tx)
}

func (e *ethProxy) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return e.eth.StorageAt(ctx, account, key, blockNumber)
}

func (e *ethProxy) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return e.eth.SubscribeFilterLogs(ctx, q, ch)
}

func (e *ethProxy) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return e.eth.SubscribeNewHead(ctx, ch)
}

func (e *ethProxy) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return e.eth.SuggestGasPrice(ctx)
}

func (e *ethProxy) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return e.eth.SyncProgress(ctx)
}

func (e *ethProxy) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return e.eth.TransactionByHash(ctx, txHash)
}

func (e *ethProxy) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return e.eth.TransactionCount(ctx, blockHash)
}

func (e *ethProxy) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	return e.eth.TransactionInBlock(ctx, blockHash, index)
}

func (e *ethProxy) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return e.eth.TransactionReceipt(ctx, txHash)
}

func (e *ethProxy) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return e.eth.SuggestGasTipCap(ctx)
}
