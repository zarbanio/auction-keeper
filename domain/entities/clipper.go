package entities

import (
	"github.com/IR-Digital-Token/auction-keeper/bindings/abacus"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Clipper struct {
	Address common.Address
	Chost   *big.Int
	Abacus  *abacus.Abacus
}
