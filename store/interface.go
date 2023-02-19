package store

import (
	"context"

	"github.com/IR-Digital-Token/auction-keeper/services/actions"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	CreateTransaction(ctx context.Context, transaction *types.Transaction, from common.Address, block uint64) (error, uint)
	GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error)
}

type ITake interface {
	CreateTake(ctx context.Context, take actions.ClipperTake, newTx types.Transaction, opts *bind.TransactOpts) (int64, error)
	// GetTakeById(ctx context.Context, id int64) (error, error)
	// GetFrobsByOrder(ctx context.Context, cursor, limit int64) ([]vat.VatFrob, error)
	// GetLastFrob(ctx context.Context) (*vat.VatFrob, error)
}
