package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type QuoteExactInputSingleParams struct {
	TokenIn           common.Address `json:"tokenIn"`
	TokenOut          common.Address `json:"tokenOut"`
	AmountIn          *big.Int       `json:"amountIn"`
	Fee               *big.Int       `json:"fee"`
	SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
}
