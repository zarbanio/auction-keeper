package callbacks

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vow"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/pubsub"
	"github.com/ethereum/go-ethereum/core/types"
)

func VowFessCallback(pubsub pubsub.Pubsub) events.CallbackFn[vow.VowFess] {
	ps := pubsub
	return func(header types.Header, fess vow.VowFess) error {
		msg := messages.NewMessage([]byte(``))
		msg.Metadata.Set("era", fess.Tab.String())

		return ps.Publish(context.Background(), "events.vow.fess", msg)
	}
}

func VowFlogCallback(pubsub pubsub.Pubsub) events.CallbackFn[vow.VowFlog] {
	ps := pubsub
	return func(header types.Header, flog vow.VowFlog) error {
		msg := messages.NewMessage([]byte(``))
		msg.Metadata.Set("era", flog.Era.String())

		return ps.Publish(context.Background(), "events.vow.flog", msg)
	}
}
