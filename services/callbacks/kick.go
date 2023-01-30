package callbacks

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
)

func ClipperKickCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperKick] {
	return func(header types.Header, kick clipper.ClipperKick) error {

		liquidatorProcessor.AddAuction(entities.Auction{
			Id:  kick.Id,
			Top: kick.Top,
			Tab: kick.Tab,
			Lot: kick.Lot,
			Usr: kick.Usr,
			Tic: header.Time,
		}, clipperName)

		return nil
	}
}
