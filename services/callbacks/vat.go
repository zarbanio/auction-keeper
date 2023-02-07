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
		msg.Metadata.Set("urnU", frob.U.String())
		msg.Metadata.Set("urnV", frob.V.String())
		msg.Metadata.Set("urnW", frob.W.String())

		return ps.Publish(context.Background(), "events.vat.frobs", msg)
	}
}
