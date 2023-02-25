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
	IClipper
	IVat
	IDog
}

type Migrateable interface {
	Migrate(path string) error
}

type ITranasaction interface {
	CreateTransaction(ctx context.Context, transaction *types.Transaction, from common.Address) (error, uint64)
	UpdateTransactionBlock(ctx context.Context, id uint64, recipt *types.Receipt, blockTime uint64, blockNumber big.Int, blockHash common.Hash) error
	GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error)
}

type IClipper interface {
	CreateTake(ctx context.Context, take *entities.ClipperTake, tx_id int64) (int64, error)
	CreateRedo(ctx context.Context, redo entities.ClipperRedo, tx_id int64) (int64, error)
}
type IVat interface {
	CreateHope(ctx context.Context, hope entities.VatHope, tx_id int64) (int64, error)
}

type IDog interface {
	CreateBark(ctx context.Context, bark entities.DogBark, tx_id int64) (int64, error)
}
