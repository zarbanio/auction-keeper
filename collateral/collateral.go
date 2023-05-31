package collateral

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
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
