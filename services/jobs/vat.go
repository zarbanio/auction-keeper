package jobs

import (
	"context"
	"github.com/IR-Digital-Token/x/messages"

	"github.com/IR-Digital-Token/auction-keeper/cache"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/ethereum/go-ethereum/common"
)

func Frobs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	urn := common.HexToAddress(msg.Metadata.Get("urn"))

	return UrnManipulation(redisCache, loader, msg.Metadata.Get("ilkName"), urn)
}

func Forks(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	urnSrc := common.HexToAddress(msg.Metadata.Get("urnSrc"))
	urnDst := common.HexToAddress(msg.Metadata.Get("urnDst"))

	return UrnManipulation(redisCache, loader, msg.Metadata.Get("ilkName"), urnSrc, urnDst)
}

func Grabs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	urn := common.HexToAddress(msg.Metadata.Get("urn"))

	return UrnManipulation(redisCache, loader, msg.Metadata.Get("ilkName"), urn)
}

func UrnManipulation(redisCache cache.ICache, loader *loaders.VaultLoader, ilkName string, urns ...common.Address) error {
	var ilkId [32]byte
	copy(ilkId[:], ilkName)

	// fetch urns from vat and update cache
	for _, urn := range urns {
		vault, err := loader.GetVaultByIlkUrn(context.Background(), ilkId, urn)
		if err != nil {
			return err
		}
		err = redisCache.SaveVault(context.Background(), *vault, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
