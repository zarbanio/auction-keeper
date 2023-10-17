package domain

import "github.com/zarbanio/auction-keeper/x/decimal"

type MarketState struct {
	TotalMarketSize map[Symbol]decimal.Decimal `json:"total_market_size"`
	TotalAvailable  map[Symbol]decimal.Decimal `json:"total_available"`
	TotalBorrows    map[Symbol]decimal.Decimal `json:"total_borrows"`
}

func (m *MarketState) AddMarketSize(symbol Symbol, size decimal.Decimal) {
	_, ok := m.TotalMarketSize[symbol]
	if !ok {
		m.TotalMarketSize[symbol] = decimal.Zero
	}
	m.TotalMarketSize[symbol] = m.TotalMarketSize[symbol].Add(size)
}

func (m *MarketState) AddTotalAvailable(symbol Symbol, size decimal.Decimal) {
	_, ok := m.TotalAvailable[symbol]
	if !ok {
		m.TotalAvailable[symbol] = decimal.Zero
	}
	m.TotalAvailable[symbol] = m.TotalAvailable[symbol].Add(size)
}

func (m *MarketState) AddTotalBorrows(symbol Symbol, size decimal.Decimal) {
	_, ok := m.TotalBorrows[symbol]
	if !ok {
		m.TotalBorrows[symbol] = decimal.Zero
	}
	m.TotalBorrows[symbol] = m.TotalBorrows[symbol].Add(size)
}
