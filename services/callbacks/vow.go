package callbacks

import (
	"context"
	"encoding/json"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vow"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/pubsub"
	"github.com/ethereum/go-ethereum/core/types"
)

func VowFessCallback(pubsub pubsub.Pubsub) events.CallbackFn[vow.VowFess] {
	ps := pubsub
	return func(header types.Header, fess vow.VowFess) error {
		payload, err := json.Marshal(fess)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.fess", messages.NewMessage(payload))
	}
}

func VowFlogCallback(pubsub pubsub.Pubsub) events.CallbackFn[vow.VowFlog] {
	ps := pubsub
	return func(header types.Header, flog vow.VowFlog) error {
		payload, err := json.Marshal(flog)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vow.flog", messages.NewMessage(payload))
	}
}
