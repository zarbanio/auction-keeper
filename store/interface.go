package store

import (
	"context"
	"math/big"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type IStore interface {
	Migrateable
	ITranasaction
	ITake
}

type Migrateable interface {
	Migrate(path string) error
}

type ITranasaction interface {
	CreateTransaction(ctx context.Context, transaction *types.Transaction, from common.Address) (error, uint64)
	UpdateTransactionBlock(ctx context.Context, id uint64, recipt *types.Receipt, blockTime uint64, blockNumber big.Int, blockHash common.Hash) error
	GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error)
}

type ITake interface {
	CreateTake(ctx context.Context, take *entities.ClipperTake, tx_id int64) (int64, error)
	// GetTakeById(ctx context.Context, id int64) (error, error)
	// GetFrobsByOrder(ctx context.Context, cursor, limit int64) ([]vat.VatFrob, error)
	// GetLastFrob(ctx context.Context) (*vat.VatFrob, error)
}
