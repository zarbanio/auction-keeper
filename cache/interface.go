package cache

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type ICache interface {
	IVault
	IIlk
}

type IIlk interface {
	SaveIlk(ctx context.Context, ilk entities.Ilk) error
	GetIlkById(ctx context.Context, id [32]byte) (*entities.Ilk, error)
}

type IVault interface {
	SaveVault(ctx context.Context, vault entities.Vault, expiration time.Duration) error
	GetVaultByIlkUrn(ctx context.Context, ilk [32]byte, urn common.Address) (*entities.Vault, error)
	getVaultByKey(ctx context.Context, key string) (*entities.Vault, error)
	GetVaults(ctx context.Context) ([]entities.Vault, error)
}
