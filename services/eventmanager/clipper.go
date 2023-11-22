package eventmanager

import (
	"math/big"

	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/services/processor"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
)

func (e *EventManager) ClipperKickCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperKick] {
	return func(raw eth.Log, kick clipper.ClipperKick) error {
		// if eventPtr > uint64(header.Number.Int64()) {
		// 	return nil
		// }
		liquidatorProcessor.AddAuction(entities.Auction{
			Id:  kick.Id,
			Top: kick.Top,
			Tab: kick.Tab,
			Lot: kick.Lot,
			Usr: kick.Usr,
			Tic: raw.Timestamp, // ?
		}, clipperName)

		return nil
	}
}

func (e *EventManager) ClipperRedoCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperRedo] {
	return func(raw eth.Log, redo clipper.ClipperRedo) error {
		// if eventPtr > uint64(header.Number.Int64()) {
		// 	return nil
		// }
		liquidatorProcessor.UpdateAuctionAfterRedo(redo.Id, redo.Top, raw.Timestamp, clipperName)
		return nil
	}
}

func (e *EventManager) ClipperTakeCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperTake] {
	return func(raw eth.Log, take clipper.ClipperTake) error {
		// if eventPtr > uint64(header.Number.Int64()) {
		// 	return nil
		// }
		if take.Lot.Cmp(big.NewInt(0)) == 0 || take.Tab.Cmp(big.NewInt(0)) == 0 {
			liquidatorProcessor.DeleteAuction(take.Id, clipperName)
		} else {
			liquidatorProcessor.UpdateAuctionAfterTake(take.Id, take.Tab, take.Lot, clipperName)
		}
		return nil
	}
}
