package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Blockchain interface {
	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
}

type SimulatedBlockchain interface {
	Commit() common.Hash
	Rollback()
	Fork(ctx context.Context, parent common.Hash) error
}

type Ethereum interface {
	ethereum.ChainReader
	ethereum.LogFilterer
	ethereum.ChainStateReader
	ethereum.ChainSyncReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.PendingContractCaller
	ethereum.PendingStateReader
	ethereum.TransactionReader
	ethereum.TransactionSender
	Blockchain
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
}

// CachedEthereum extends the Ethereum interface with a method for cached log filtering.
type CachedEthereum interface {
	Ethereum // Embedding the Ethereum interface.
	// CachedFilterLogs attempts to retrieve logs based on the provided filter query.
	// If the logs corresponding to the query are found in the cache, they are returned.
	// Otherwise, the method queries the Ethereum network and updates the cache with the
	// retrieved logs before returning them. An error is returned if the log retrieval fails.
	CachedFilterLogs(context context.Context, q ethereum.FilterQuery) ([]Log, error)

	// CachedSubscribeFilterLogs subscribes to new log entries that match the provided filter query,
	// with an added caching behavior to improve performance.
	CachedSubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- Log) (ethereum.Subscription, error)
}

// Log represents an entry of logging information with an identifier and timestamp.
// It embeds the types.Log struct to inherit fields relevant to Ethereum logs.
type Log struct {
	Id        uint64    `json:"id"`        // Unique identifier for the log entry.
	Timestamp time.Time `json:"timestamp"` // Timestamp of when the log entry was created.
	types.Log           // Embedding the types.Log struct to provide Ethereum log fields.
}
