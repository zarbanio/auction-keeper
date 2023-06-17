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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/events"
)

type ReceiptCallbackFunc func(*types.Receipt, *types.Header) error

type Indexer struct {
	eth           *ethclient.Client
	bcache        *BlockCache
	addresses     []common.Address
	eventHandlers map[string]events.Handler
	blockInterval time.Duration
	blockPointer  BlockPointer
}

func NewIndexer(
	eth *ethclient.Client,
	bcache *BlockCache,
	blockInterval time.Duration,
	blockPtr BlockPointer,
	addresses []common.Address,
	eventHandlers map[string]events.Handler) *Indexer {

	return &Indexer{
		eth:           eth,
		bcache:        bcache,
		addresses:     addresses,
		eventHandlers: eventHandlers,
		blockInterval: blockInterval,
		blockPointer:  blockPtr,
	}
}

func (i *Indexer) IndexHistory(start, latestBlock *big.Int) error {
	blockRange := int64(2000)

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

		events, err := i.eth.FilterLogs(context.Background(), query)
		if err != nil {
			return err
		}

		for _, event := range events {
			if len(event.Topics) == 0 {
				continue
			}
			handler, ok := i.eventHandlers[event.Topics[0].String()]
			if !ok {
				continue
			}
			block, err := i.bcache.GetBlockByNumber(context.Background(), event.BlockNumber)
			if err != nil {
				return err
			}
			err = handler.DecodeAndHandle(*block.Header(), event)
			if err != nil {
				return err
			}
		}
		err = i.blockPointer.Update(toBlock.Uint64())
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Indexer) IndexNewEvents() error {
	i.trackBlockPtr()

	newEvents := make(chan types.Log)
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: i.addresses,
	}
	sub, err := i.eth.SubscribeFilterLogs(context.Background(), query, newEvents)
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
			handler, ok := i.eventHandlers[event.Topics[0].String()]
			if !ok {
				continue
			}
			block, err := i.bcache.GetBlockByNumber(context.Background(), event.BlockNumber)
			if err != nil {
				return err
			}
			err = handler.DecodeAndHandle(*block.Header(), event)
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
	_, err := i.eth.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("error subscribing to new heads.", err)
	}
	go func() {
		for header := range headers {
			err = i.blockPointer.Update(header.Number.Uint64())
			if err != nil {
				log.Fatal("error updating block pointer", err)
			}
		}
	}()
}

func (i *Indexer) WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, *types.Header, error) {
	ticker := time.NewTicker(i.blockInterval)
	defer ticker.Stop()

	var receipt *types.Receipt
	var err error
	for receipt == nil {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-ticker.C:
			receipt, err = i.eth.TransactionReceipt(ctx, txHash)
			if receipt != nil {
				break
			}
		}
	}

	block, err := i.bcache.GetBlockByNumber(ctx, receipt.BlockNumber.Uint64())
	if err != nil {
		return nil, nil, err
	}
	return receipt, block.Header(), nil
}

func (i *Indexer) SubmitTxAndCallOnReceipt(ctx context.Context, tx *types.Transaction, callback ReceiptCallbackFunc) error {
	err := i.eth.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("error sending transaction. %w", err)
	}

	receipt, header, err := i.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("error waiting for receipt. %w", err)
	}

	return callback(receipt, header)
}
