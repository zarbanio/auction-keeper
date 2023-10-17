package domain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type ProtocolData struct {
	ReservesData     []FormattedReserveData `json:"reserves_data"`
	BaseCurrencyData BaseCurrencyData       `json:"base_currency_data"`
	MarketState      MarketState            `json:"market_state"`
}

func GetAssetSymbol(p ProtocolData, addr common.Address) Symbol {
	for _, r := range p.ReservesData {
		if r.UnderlyingAsset.String() == addr.String() {
			return Symbol(r.Symbol)
		}
	}
	return ""
}

func GetFormattedReserveData(p ProtocolData, addr common.Address) *FormattedReserveData {
	for _, r := range p.ReservesData {
		if r.UnderlyingAsset.String() == addr.String() {
			return &r
		}
	}
	return nil
}

func GetAssetPrices(p ProtocolData, addr common.Address) decimal.Decimal {
	for _, r := range p.ReservesData {
		if r.UnderlyingAsset.String() == addr.String() {
			return r.Price
		}
	}
	return decimal.Zero
}

func GetAssetSupplyAPY(p ProtocolData, addr common.Address) decimal.Decimal {
	for _, r := range p.ReservesData {
		if r.UnderlyingAsset.String() == addr.String() {
			return r.SupplyAPY
		}
	}
	return decimal.Zero
}

func GetAssetVariableBorrowAPY(p ProtocolData, addr common.Address) decimal.Decimal {
	for _, r := range p.ReservesData {
		if r.UnderlyingAsset.String() == addr.String() {
			return r.VariableBorrowAPY
		}
	}
	return decimal.Zero
}
