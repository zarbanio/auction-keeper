// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package lendingpool

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type FlashLoanHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolFlashLoan]
}

func (h *FlashLoanHandler) ID() string {
	return h.addr + ":" + "0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac"
}

func (h *FlashLoanHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlashLoan(log.Log)
}

func (h *FlashLoanHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolFlashLoan)
	if !ok {
		return errors.New("event type is not LendingpoolFlashLoan")
	}
	return h.callback(log, e)
}

func (h *FlashLoanHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlashLoan(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFlashLoanHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolFlashLoan]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlashLoanHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
