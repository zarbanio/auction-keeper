package collateral

import (
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/ethereum/go-ethereum/common"
)

type Collateral struct {
	Name                 string
	Clipper              entities.Clipper
	ClipperLoader        *loaders.ClipperLoader
	GemJoinAdapter       common.Address
	UniswapV3Callee      common.Address
	UniswapV3CalleeRoute []byte
}
