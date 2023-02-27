package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ClipperTake struct {
	Auction_id *big.Int
	Amt        *big.Int
	Max        *big.Int
	Who        common.Address
	Data       []byte
}
