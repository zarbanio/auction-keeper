package callbacks

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

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
