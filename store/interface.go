package store

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	StoreTakeTransaction(ctx context.Context, newTx types.Transaction, opts *bind.TransactOpts, id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) (int64, error)
	// GetTakeById(ctx context.Context, id int64) (error, error)
	// GetFrobsByOrder(ctx context.Context, cursor, limit int64) ([]vat.VatFrob, error)
	// GetLastFrob(ctx context.Context) (*vat.VatFrob, error)
}
