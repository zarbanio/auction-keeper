package store

import (
	"context"

	"github.com/IR-Digital-Token/auction-keeper/services/actions"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type IStore interface {
	Migrateable
	ITake
}

type Migrateable interface {
	Migrate(path string) error
}

type ITake interface {
	CreateTake(ctx context.Context, take actions.ClipperTake, newTx types.Transaction, opts *bind.TransactOpts) (int64, error)
	// GetTakeById(ctx context.Context, id int64) (error, error)
	// GetFrobsByOrder(ctx context.Context, cursor, limit int64) ([]vat.VatFrob, error)
	// GetLastFrob(ctx context.Context) (*vat.VatFrob, error)
}
