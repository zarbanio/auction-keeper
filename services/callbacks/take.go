package callbacks

import (
	"fmt"
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
)

func ClipperTakeCallback() events.CallbackFn[clipper.ClipperTake] {
	return func(header types.Header, take clipper.ClipperTake) error {
		fmt.Println("__Take")
		fmt.Println("\tID", take.Id)
		return nil
	}
}
