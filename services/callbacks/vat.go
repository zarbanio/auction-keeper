package callbacks

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/pubsub"
	"github.com/ethereum/go-ethereum/core/types"
)

func VatFrobCallback(pubsub pubsub.Pubsub) events.CallbackFn[vat.VatFrob] {
	ps := pubsub
	return func(header types.Header, frob vat.VatFrob) error {
		msg := messages.NewMessage([]byte(``))
		msg.Metadata.Set("ilkName", string(frob.I[:]))
		msg.Metadata.Set("urn", frob.U.String())

		return ps.Publish(context.Background(), "events.vat.frobs", msg)
	}
}

func VatForkCallback(pubsub pubsub.Pubsub) events.CallbackFn[vat.VatFork] {
	ps := pubsub
	return func(header types.Header, fork vat.VatFork) error {
		msg := messages.NewMessage([]byte(``))
		msg.Metadata.Set("ilkName", string(fork.Ilk[:]))
		msg.Metadata.Set("urnSrc", fork.Src.String())
		msg.Metadata.Set("urnDst", fork.Dst.String())

		return ps.Publish(context.Background(), "events.vat.forks", msg)
	}
}

func VatGrabCallback(pubsub pubsub.Pubsub) events.CallbackFn[vat.VatGrab] {
	ps := pubsub
	return func(header types.Header, grab vat.VatGrab) error {
		msg := messages.NewMessage([]byte(``))
		msg.Metadata.Set("ilkName", string(grab.I[:]))
		msg.Metadata.Set("urn", grab.U.String())

		return ps.Publish(context.Background(), "events.vat.grabs", msg)
	}
}
