package collateral

import (
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/uniswap_v3"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Collateral struct {
	Name            string
	Decimal         *big.Int
	Clipper         entities.Clipper
	ClipperLoader   *loaders.ClipperLoader
	GemJoinAdapter  common.Address
	UniswapV3Callee common.Address
	UniswapV3Path   []entities.UniswapV3Route
	UniswapV3Quoter *uniswap_v3.UniswapV3Quoter
}
