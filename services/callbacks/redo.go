package callbacks

import (
	"fmt"
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/core/types"
)

func ClipperRedoCallback() events.CallbackFn[clipper.ClipperRedo] {
	return func(header types.Header, redo clipper.ClipperRedo) error {
		fmt.Println("__Redo")
		fmt.Println("\tID", redo.Id)
		return nil
	}
}
