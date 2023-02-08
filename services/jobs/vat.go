package jobs

import (
	"context"
	"errors"
	"github.com/IR-Digital-Token/x/messages"

	"github.com/IR-Digital-Token/auction-keeper/cache"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/ethereum/go-ethereum/common"
)

func Frobs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	var ilkId [32]byte
	copy(ilkId[:], msg.Metadata.Get("ilkName"))
	urn := common.HexToAddress(msg.Metadata.Get("urn"))

	return UrnManipulation(redisCache, loader, ilkId, urn)
}

func Forks(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	var ilkId [32]byte
	copy(ilkId[:], msg.Metadata.Get("ilkName"))
	urnSrc := common.HexToAddress(msg.Metadata.Get("urnSrc"))
	urnDst := common.HexToAddress(msg.Metadata.Get("urnDst"))

	return UrnManipulation(redisCache, loader, ilkId, urnSrc, urnDst)
}

func Grabs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	var ilkId [32]byte
	copy(ilkId[:], msg.Metadata.Get("ilkName"))
	urn := common.HexToAddress(msg.Metadata.Get("urn"))

	return UrnManipulation(redisCache, loader, ilkId, urn)
}

func UrnManipulation(redisCache cache.ICache, loader *loaders.VaultLoader, ilkId [32]byte, urns ...common.Address) error {
	// try to get ilk from redis. if it doesn't exist, fetch it from vat and cache it
	ilk, err := redisCache.GetIlkById(context.Background(), ilkId)
	if errors.Is(err, cache.ErrIlkNotFound) {
		ilk, err = loader.GetIlkById(context.Background(), ilkId)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// fetch urns from vat and update cache
	for _, urn := range urns {
		vault, err := loader.GetVaultByIlkUrn(context.Background(), ilk, urn)
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
