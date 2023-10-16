package store

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/bindings/zarban/median"
	"github.com/zarbanio/auction-keeper/domain"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type IStore interface {
	Migrateable
	ITranasaction
	IClipper
	IVat
	IDog
	IVow
	IBlockPtr
	IEthereumLog
	ILogMedianPrice
	IIlk
}

type Migrateable interface {
	Migrate(path string) error
}

type ITranasaction interface {
	CreateTransaction(ctx context.Context, transaction *types.Transaction) (error, uint64)
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

type IVow interface {
	CreateKiss(ctx context.Context, kiss *entities.VowKiss, tx_id int64) (int64, error)
	CreateHeal(ctx context.Context, heal *entities.VowHeal, tx_id int64) (int64, error)
	CreateFlog(ctx context.Context, flog *entities.VowFlog, tx_id int64) (int64, error)
	CreateFlop(ctx context.Context, tx_id int64) (int64, error)
}

type IBlockPtr interface {
	ExistsBlockPtr(ctx context.Context) (bool, error)
	CreateBlockPtr(ctx context.Context, ptr uint64) (int64, error)
	GetBlockPtrById(ctx context.Context, id int64) (uint64, error)
	GetLastBlockPtr(ctx context.Context) (int64, uint64, error)
	UpdateBlockPtr(ctx context.Context, id int64, ptr uint64) (uint64, error)
	IncBlockPtr(ctx context.Context, id int64) (uint64, error)
	DeleteBlockPtr(ctx context.Context, id int64) error
}

type EthereumLogsQuery struct {
	Id        int64
	FromBlock uint64
	ToBlock   uint64
	Addresses []common.Address
}

type IEthereumLog interface {
	CreateLog(ctx context.Context, log eth.Log) (uint64, error)
	GetLogById(ctx context.Context, id uint64) (*eth.Log, error)
	GetLogsByQuery(ctx context.Context, q EthereumLogsQuery) ([]eth.Log, error)
}

type ILogMedianPrice interface {
	CreateLogMedianPrice(ctx context.Context, logMedainPrice median.MedianLogMedianPrice, logId uint64) (int64, error)
	GetLogMedianPriceById(ctx context.Context, id int64) (*median.MedianLogMedianPrice, error)
	GetLogMedianPriceByOrder(ctx context.Context, address common.Address, cursor, limit int64) ([]median.MedianLogMedianPrice, error)
	GetLastLogMedianPrice(ctx context.Context, address common.Address) (*median.MedianLogMedianPrice, error)
}

type IIlk interface {
	CreateOrUpdateIlk(ctx context.Context, ilk domain.Ilk) error
}
