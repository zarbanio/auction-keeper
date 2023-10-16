package chain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
)

type ReceiptCallbackFunc func(*types.Receipt) error

type Indexer struct {
	ceth             eth.CachedEthereum
	addresses        []common.Address
	eventHandlers    map[string]events.Handler
	blockInterval    time.Duration
	liveBlockPointer BlockPointer
	historyBlockPtr  BlockPointer
	blockRange       int64
}

func NewIndexer(
	ceth eth.CachedEthereum,
	blockInterval time.Duration,
	liveBlockPointer BlockPointer,
	historyBlockPtr BlockPointer,
	blockRange int64,
	addresses []common.Address,
	eventHandlers map[string]events.Handler) *Indexer {

	return &Indexer{
		ceth:             ceth,
		addresses:        addresses,
		eventHandlers:    eventHandlers,
		blockInterval:    blockInterval,
		liveBlockPointer: liveBlockPointer,
		historyBlockPtr:  historyBlockPtr,
		blockRange:       blockRange,
	}
}

func (i *Indexer) IndexHistory(start, latestBlock *big.Int) error {
	blockRange := i.blockRange

	for fromBlock := start; fromBlock.Cmp(latestBlock) < 0; fromBlock = new(big.Int).Add(fromBlock, big.NewInt(blockRange)) {
		toBlock := new(big.Int).Add(fromBlock, big.NewInt(int64(blockRange-1)))
		if toBlock.Cmp(latestBlock) > 0 {
			toBlock = latestBlock
		}

		query := ethereum.FilterQuery{
			FromBlock: fromBlock,
			ToBlock:   toBlock,
			Addresses: i.addresses,
		}

		events, err := i.ceth.CachedFilterLogs(context.Background(), query)
		if err != nil {
			return err
		}

		for _, event := range events {
			if len(event.Topics) == 0 {
				continue
			}
			id := event.Address.String() + ":" + event.Topics[0].String()
			handler, ok := i.eventHandlers[id]
			if !ok {
				continue
			}
			err = handler.DecodeAndHandle(event)
			if err != nil {
				return err
			}
		}
		err = i.historyBlockPtr.Update(toBlock.Uint64())
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Indexer) IndexNewEvents() error {
	i.trackBlockPtr()

	newEvents := make(chan eth.Log)
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: i.addresses,
	}
	sub, err := i.ceth.CachedSubscribeFilterLogs(context.Background(), query, newEvents)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-newEvents:
			if len(event.Topics) == 0 {
				continue
			}
			id := event.Address.String() + ":" + event.Topics[0].String()
			handler, ok := i.eventHandlers[id]
			if !ok {
				continue
			}
			err = handler.DecodeAndHandle(event)
			if err != nil {
				return err
			}
		case err := <-sub.Err():
			return err
		}
	}
}

func (i *Indexer) trackBlockPtr() {
	headers := make(chan *types.Header)
	_, err := i.ceth.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("error subscribing to new heads.", err)
	}
	go func() {
		for header := range headers {
			err = i.liveBlockPointer.Update(header.Number.Uint64())
			if err != nil {
				log.Fatal("error updating block pointer", err)
			}
		}
	}()
}

func (i *Indexer) WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ticker := time.NewTicker(i.blockInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			receipt, err := i.ceth.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
		}
	}
}

func (i *Indexer) SubmitTxAndCallOnReceipt(ctx context.Context, tx *types.Transaction, callback ReceiptCallbackFunc) error {
	err := i.ceth.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("error sending transaction. %w", err)
	}

	receipt, err := i.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("error waiting for receipt. %w", err)
	}

	return callback(receipt)
}
