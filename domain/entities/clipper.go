package entities

import (
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/bindings/abacus"
	"github.com/ethereum/go-ethereum/common"
)

type Clipper struct {
	Address common.Address
	Chost   *big.Int
	Abacus  *abacus.Abacus
}
