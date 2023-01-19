package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type Collateral struct {
	Name           string
	Clipper        Clipper
	GemJoinAdapter common.Address
}
