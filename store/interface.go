package store

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/bindings/zarban/median"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
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
	IEthereumLogsCacheKey
	IFrob
	IVatGrab
	IFork
	IFlog
	IFess
	IFlop
	ILog
}

type Migrateable interface {
	Migrate(path string) error
}

type ITranasaction interface {
	CreateTransaction(ctx context.Context, transaction *types.Transaction) (error, uint64)
	UpdateTransactionReceipt(ctx context.Context, id uint64, recipt *types.Receipt, header *types.Header) error
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
}

type IBlockPtr interface {
	ExistsBlockPtr(ctx context.Context, name string) (bool, error)
	CreateBlockPtr(ctx context.Context, name string, ptr uint64) (int64, error)
	GetBlockPtrByName(ctx context.Context, name string) (uint64, error)
	UpdateBlockPtr(ctx context.Context, name string, ptr uint64) (uint64, error)
	IncBlockPtr(ctx context.Context, name string) (uint64, error)
	DeleteBlockPtr(ctx context.Context, name string) error
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
	GetIlkByName(ctx context.Context, name string) (*domain.Ilk, error)
}

type IEthereumLogsCacheKey interface {
	CreateEthereumLogsCacheKey(ctx context.Context, key EthereumLogsQuery) error
	IsEthereumLogsCached(ctx context.Context, key EthereumLogsQuery) (bool, error)
}

type IFrob interface {
	CreateFrob(ctx context.Context, frob vat.VatFrob, logId uint64) (int64, error)
}

type IVatGrab interface {
	CreateVatGrab(ctx context.Context, vatGrab vat.VatGrab, logId uint64) (int64, error)
}

type IFork interface {
	CreateFork(ctx context.Context, move vat.VatFork, logId uint64) (int64, error)
}

type IFlog interface {
	CreateFlog(ctx context.Context, flog vow.VowFlog, logId uint64) (int64, error)
}

type IFess interface {
	CreateFess(ctx context.Context, fess vow.VowFess, logId uint64) (int64, error)
}

type IFlop interface {
	CreateFlop(ctx context.Context, tx_id int64) (int64, error)
}

type ILog interface {
	CreateLogs(ctx context.Context, b []byte) (int, error)
}
