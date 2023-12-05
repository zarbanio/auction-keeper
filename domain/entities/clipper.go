package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/bindings/zarban/abacus"
)

type Clipper struct {
	Address common.Address
	Chost   *big.Int
	Abacus  *abacus.Abacus
}
