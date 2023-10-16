package eventmanager

import (
	"context"
	"encoding/json"

	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/zarbanio/auction-keeper/x/messages"
	"github.com/zarbanio/auction-keeper/x/pubsub"
)

func VowFessCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vow.VowFess] {
	ps := pubsub
	return func(raw eth.Log, fess vow.VowFess) error {
		// if eventPtr > uint64(header.Number.Int64()) {
		// 	return nil
		// }
		payload, err := json.Marshal(fess)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.fess", messages.NewMessage(payload))
	}
}

func VowFlogCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vow.VowFlog] {
	ps := pubsub
	return func(raw eth.Log, flog vow.VowFlog) error {
		// if eventPtr > uint64(header.Number.Int64()) {
		// 	return nil
		// }
		payload, err := json.Marshal(flog)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.flog", messages.NewMessage(payload))
	}
}
