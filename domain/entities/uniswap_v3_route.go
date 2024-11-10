package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type UniswapV3Route struct {
	Fee    int64
	TokenA common.Address
	TokenB common.Address
}
