package callbacks

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/bindings/vow"
	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/zarbanio/auction-keeper/x/messages"
	"github.com/zarbanio/auction-keeper/x/pubsub"
)

func VowFessCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vow.VowFess] {
	ps := pubsub
	return func(header types.Header, fess vow.VowFess) error {
		if eventPtr > uint64(header.Number.Int64()) {
			return nil
		}
		payload, err := json.Marshal(fess)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.fess", messages.NewMessage(payload))
	}
}

func VowFlogCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vow.VowFlog] {
	ps := pubsub
	return func(header types.Header, flog vow.VowFlog) error {
		if eventPtr > uint64(header.Number.Int64()) {
			return nil
		}
		payload, err := json.Marshal(flog)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.flog", messages.NewMessage(payload))
	}
}
