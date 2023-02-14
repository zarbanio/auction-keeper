package cache

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type ICache interface {
	IVatIlk
	IVault
}

type IVatIlk interface {
	SaveVatIlk(ctx context.Context, ilk entities.VatIlk, expiration time.Duration) error
	GetVatIlkById(ctx context.Context, id [32]byte) (*entities.VatIlk, error)
}

type IVault interface {
	SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error
	GetVaults(ctx context.Context) ([]*entities.Vault, error)
	GetVaultByIlkUrn(ctx context.Context, ilk [32]byte, urn common.Address) (*entities.Vault, error)
	getVaultByKey(ctx context.Context, key string) (*entities.Vault, error)
}
