package eoa

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/zarbanio/auction-keeper/bindings/ierc20"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type EOA struct {
	sender sender.Sender
	erc20s map[string]*ierc20.Ierc20
}

func NewEOA(sender sender.Sender, erc20s map[string]*ierc20.Ierc20) *EOA {
	return &EOA{
		sender: sender,
		erc20s: erc20s,
	}
}

func (e *EOA) GetBalance() (map[string]decimal.Decimal, error) {
	balances := make(map[string]decimal.Decimal)
	for symbol, erc20 := range e.erc20s {
		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: context.Background()}, e.sender.GetAddress())
		if err != nil {
			return nil, err
		}
		decBalance := decimal.NewFromBigInt(balance)
		decBalance = decBalance.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)))
		balances[symbol] = decBalance
	}
	return balances, nil
}
