package callbacks

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
)

func ClipperKickCallback(auctionProcessor *processor.AuctionProcessor, clipperName string) events.CallbackFn[clipper.ClipperKick] {
	return func(header types.Header, kick clipper.ClipperKick) error {

		auctionProcessor.AddAuction(entities.Auction{
			Id:  decimal.NewFromBigInt(kick.Id, 0),
			Top: decimal.NewFromBigInt(kick.Top, 0),
			Tab: decimal.NewFromBigInt(kick.Tab, 0),
			Lot: decimal.NewFromBigInt(kick.Lot, 0),
			Usr: kick.Usr,
			Tic: header.Time,

			ClipperName: clipperName,
		})

		return nil
	}
}
