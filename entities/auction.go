package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type Auction struct {
	Id  decimal.Decimal
	Top decimal.Decimal // Initial price of collateral in the auction (with the current OSM price and buf percent)
	Tab decimal.Decimal // Debt: the target SIM to raise from the auction (debt + stability fees + liquidation penalty) [rad]
	Lot decimal.Decimal // Collateral: the amount of collateral available for purchase [wad]
	Usr common.Address  // Address that will receive any leftover collateral
	Tic uint64          // Auction start time [unix epoch]

	ClipperName string
}
