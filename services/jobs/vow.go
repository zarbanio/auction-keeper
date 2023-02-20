package jobs

import (
	"context"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/cache"
	"github.com/IR-Digital-Token/x/messages"
	"math/big"
)

func Fess(msg *messages.Message, redisCache cache.ICache) error {
	msgValue := msg.Metadata.Get("era")

	era, success := new(big.Int).SetString(msgValue, 10)
	if !success {
		return fmt.Errorf("error in convert era in msg to big.int. msg value: %s", msgValue)
	}

	return redisCache.SaveEra(context.Background(), era, 0)
}

func Flog(msg *messages.Message, redisCache cache.ICache) error {
	msgValue := msg.Metadata.Get("era")

	era, success := new(big.Int).SetString(msgValue, 10)
	if !success {
		return fmt.Errorf("error in convert era in msg to big.int. msg value: %s", msgValue)
	}

	return redisCache.DeleteEra(context.Background(), era)
}
