package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/x/messages"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/services/loaders"
)

func Frobs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {

	var frob vat.VatFrob
	if err := json.Unmarshal(msg.Payload, &frob); err != nil {
		return fmt.Errorf("error in unmarshalling: %v", err)
	}

	return UrnManipulation(redisCache, loader, frob.I, frob.U)
}

func Forks(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	var fork vat.VatFork
	if err := json.Unmarshal(msg.Payload, &fork); err != nil {
		return fmt.Errorf("error in unmarshalling: %v", err)
	}
	return UrnManipulation(redisCache, loader, fork.Ilk, fork.Src, fork.Dst)
}

func Grabs(msg *messages.Message, redisCache cache.ICache, loader *loaders.VaultLoader) error {
	var grab vat.VatGrab
	if err := json.Unmarshal(msg.Payload, &grab); err != nil {
		return fmt.Errorf("error in unmarshalling: %v", err)
	}
	return UrnManipulation(redisCache, loader, grab.I, grab.U)
}

func UrnManipulation(redisCache cache.ICache, loader *loaders.VaultLoader, ilkId [32]byte, urns ...common.Address) error {
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
