package cache

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type ICache interface {
	IVatIlk
	IVault
	IEra
}

type MemCache interface {
	IAddress
}

type IVatIlk interface {
	SaveVatIlk(ctx context.Context, ilk entities.VatIlk, expiration time.Duration) error
	GetVatIlkByName(ctx context.Context, ilkName string) (*entities.VatIlk, error)
}

type IVault interface {
	SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error
	GetVaults(ctx context.Context) ([]*entities.Vault, error)
	getVaultByKey(ctx context.Context, key string) (*entities.Vault, error)
}

type IEra interface {
	SaveEra(ctx context.Context, era *big.Int, expiration time.Duration) error
	DeleteEra(ctx context.Context, era *big.Int) error
	GetEras(ctx context.Context) ([]*big.Int, error)
}

type IAddress interface {
	SaveAddresses(ctx context.Context, addresses map[string]common.Address) error
	GetAddresses(ctx context.Context) (map[string]common.Address, error)
}
