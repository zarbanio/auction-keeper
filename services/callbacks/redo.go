package callbacks

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
)

func ClipperRedoCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperRedo] {
	return func(header types.Header, redo clipper.ClipperRedo) error {
		liquidatorProcessor.UpdateAuctionAfterRedo(redo.Id, redo.Top, header.Time, clipperName)
		return nil
	}
}
