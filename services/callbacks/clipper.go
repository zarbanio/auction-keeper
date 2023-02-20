package callbacks

import (
	"math/big"

	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
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

func ClipperRedoCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperRedo] {
	return func(header types.Header, redo clipper.ClipperRedo) error {
		liquidatorProcessor.UpdateAuctionAfterRedo(redo.Id, redo.Top, header.Time, clipperName)
		return nil
	}
}

func ClipperTakeCallback(liquidatorProcessor *processor.LiquidatorProcessor, clipperName string) events.CallbackFn[clipper.ClipperTake] {
	return func(header types.Header, take clipper.ClipperTake) error {
		if take.Lot.Cmp(big.NewInt(0)) == 0 || take.Tab.Cmp(big.NewInt(0)) == 0 {
			liquidatorProcessor.DeleteAuction(take.Id, clipperName)
		} else {
			liquidatorProcessor.UpdateAuctionAfterTake(take.Id, take.Tab, take.Lot, clipperName)
		}
		return nil
	}
}
