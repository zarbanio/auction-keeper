package callbacks

import (
	"context"
	"encoding/json"

	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/pubsub"
	"github.com/ethereum/go-ethereum/core/types"
)

func VatFrobCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vat.VatFrob] {
	ps := pubsub
	return func(header types.Header, frob vat.VatFrob) error {
		if eventPtr > uint64(header.Number.Int64()) {
			return nil
		}
		payload, err := json.Marshal(frob)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vat.frobs", messages.NewMessage(payload))
	}
}

func VatForkCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vat.VatFork] {
	ps := pubsub
	return func(header types.Header, fork vat.VatFork) error {
		if eventPtr > uint64(header.Number.Int64()) {
			return nil
		}
		payload, err := json.Marshal(fork)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vat.forks", messages.NewMessage(payload))
	}
}

func VatGrabCallback(pubsub pubsub.Pubsub, eventPtr uint64) events.CallbackFn[vat.VatGrab] {
	ps := pubsub
	return func(header types.Header, grab vat.VatGrab) error {
		if eventPtr > uint64(header.Number.Int64()) {
			return nil
		}
		payload, err := json.Marshal(grab)
		if err != nil {
			return err
		}

		return ps.Publish(context.Background(), "events.vat.grabs", messages.NewMessage(payload))
	}
}
