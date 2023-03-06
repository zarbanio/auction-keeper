package cache

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"math/big"
	"time"
)

type ICache interface {
	IVatIlk
	IVault
	IEra
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
