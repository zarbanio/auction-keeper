package entities

import (
	"math/big"

	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Auction struct {
	Id  *big.Int
	Top *big.Int       // Initial price of collateral in the auction (with the current OSM price and buf percent)
	Tab *big.Int       // Debt: the target ZAR to raise from the auction (debt + stability fees + liquidation penalty) [rad]
	Lot *big.Int       // Collateral: the amount of collateral available for purchase [wad]
	Usr common.Address // ClipperAddress that will receive any leftover collateral
	Tic time.Time      // Auction start time [unix epoch]
}
