package chain

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/panjf2000/ants/v2"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/zarbanio/auction-keeper/x/transactions"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

type PollingIndexer struct {
	Head          chan uint64
	ptr           uint64
	batchSize     uint64
	blockInterval time.Duration
	pool          *ants.Pool
	eth           Ethereum
	blockPointer  BlockPointer
	logHandlers   map[string]events.Handler
	addresses     map[string]bool
	txWatchList   map[common.Hash]transactions.Handler
	mutex         *sync.Mutex
}

func NewPollingIndexer(eth Ethereum, blockPointer BlockPointer, poolSize int) *PollingIndexer {
	pool, err := ants.NewPool(poolSize)
	if err != nil {
		panic(err)
	}

	return &PollingIndexer{
		eth:          eth,
		blockPointer: blockPointer,
		logHandlers:  make(map[string]events.Handler),
		pool:         pool,
		batchSize:    uint64(poolSize * 10),
		addresses:    make(map[string]bool),
		txWatchList:  make(map[common.Hash]transactions.Handler),
		mutex:        &sync.Mutex{},
	}
}

func (p *PollingIndexer) Init(blockInterval time.Duration) {
	ptr, err := p.blockPointer.Read()
	if err != nil {
		panic(err)
	}
	p.ptr = ptr

	head, err := HeadChannel(p.eth, blockInterval)
	if err != nil {
		panic(err)
	}
	p.Head = head
	p.blockInterval = blockInterval
}

func (p *PollingIndexer) Start() error {
	head := <-p.Head
	for p.ptr < head {
		err := p.loop(p.ptr, min(p.ptr+p.batchSize, head))
		if err != nil {
			return err
		}

		diff := head - p.ptr
		if diff < p.batchSize {
			p.ptr += diff
		} else {
			p.ptr += p.batchSize
		}
		err = p.blockPointer.Update(p.ptr)
		if err != nil {
			return err
		}
	}
	return nil
}

func min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func (p *PollingIndexer) loop(from, to uint64) error {
	ch := make(chan error)
	done := make(chan struct{})
	parent := context.Background()
	go func() {
		wg := &sync.WaitGroup{}
		for i := from; i < to; i++ {
			wg.Add(1)
			j := i
			err := p.pool.Submit(func() {
				ctx, cancel := context.WithTimeout(parent, p.blockInterval)
				defer cancel()
				err := p.processBlock(ctx, big.NewInt(int64(j)))
				if err != nil {
					ch <- err
					return
				}
				wg.Done()
			})
			if err != nil {
				ch <- err
				return
			}
		}
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return nil
	case err := <-ch:
		return err
	}
}

func (p *PollingIndexer) processBlock(ctx context.Context, number *big.Int) error {
	block, err := p.eth.BlockByNumber(ctx, number)
	if err != nil {
		return err
	}

	logs, err := p.eth.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: block.Number(),
		ToBlock:   block.Number(),
	})
	if err != nil {
		return err
	}
	err = p.processTransactions(*block.Header(), p.filterTxHash(block.Transactions()))
	if err != nil {
		return err
	}
	return p.processLogs(*block.Header(), p.filterLogs(logs))
}

func (p *PollingIndexer) processTransactions(header types.Header, txList types.Transactions) error {
	for _, tx := range txList {
		txHash := tx.Hash()
		txRecipt, err := p.eth.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return err
		}

		handler, ok := p.txWatchList[txHash]
		if !ok {
			continue
		}
		err = handler.HandleTransaction()(header, txRecipt)
		if err != nil {
			return err
		}
		p.UnWatchTx(handler)
	}
	return nil
}

func (p *PollingIndexer) processLogs(header types.Header, logs []types.Log) error {
	for _, l := range logs {
		if len(l.Topics) == 0 {
			continue
		}
		handler, ok := p.logHandlers[l.Topics[0].String()]
		if !ok {
			continue
		}

		ethLog := eth.Log{Log: l}
		err := handler.DecodeAndHandle(ethLog)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PollingIndexer) filterLogs(logs []types.Log) []types.Log {
	var res []types.Log
	for _, l := range logs {
		_, ok := p.addresses[l.Address.String()]

		if ok {
			res = append(res, l)
		}
	}
	return res
}
func (p *PollingIndexer) filterTxHash(transactions types.Transactions) types.Transactions {
	var res types.Transactions
	for _, tx := range transactions {
		txHash := tx.Hash()
		_, ok := p.txWatchList[txHash]
		if ok {
			res = append(res, tx)
		}
	}
	return res
}

func (p *PollingIndexer) RegisterEventHandlers(handlers ...events.Handler) {
	for _, handler := range handlers {
		p.logHandlers[handler.ID()] = handler
	}
}

func (p *PollingIndexer) RegisterAddresses(addresses ...common.Address) {
	for _, addr := range addresses {
		p.addresses[addr.String()] = true
	}
}

func (p *PollingIndexer) WatchTx(handler transactions.Handler) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.txWatchList[handler.ID()] = handler
}

func (p *PollingIndexer) UnWatchTx(txHandler transactions.Handler) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.txWatchList, txHandler.ID())
}

func (p *PollingIndexer) Ptr() uint64 {
	return p.ptr
}
