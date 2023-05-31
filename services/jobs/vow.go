package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zarbanio/auction-keeper/bindings/vow"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/x/messages"
)

func Fess(msg *messages.Message, redisCache cache.ICache) error {
	var fess vow.VowFess
	if err := json.Unmarshal(msg.Payload, &fess); err != nil {
		return fmt.Errorf("error in unmarshalling era variable: %v", err)
	}

	return redisCache.SaveEra(context.Background(), fess.Tab, 0)
}

func Flog(msg *messages.Message, redisCache cache.ICache) error {
	var flog vow.VowFlog
	if err := json.Unmarshal(msg.Payload, &flog); err != nil {
		return fmt.Errorf("error in unmarshalling era variable: %v", err)
	}

	return redisCache.DeleteEra(context.Background(), flog.Era)
}
