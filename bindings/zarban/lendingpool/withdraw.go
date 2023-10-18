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

type WithdrawHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolWithdraw]
}

func (h *WithdrawHandler) ID() string {
	return h.addr + ":" + "0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7"
}

func (h *WithdrawHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseWithdraw(log.Log)
}

func (h *WithdrawHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolWithdraw)
	if !ok {
		return errors.New("event type is not LendingpoolWithdraw")
	}
	return h.callback(log, e)
}

func (h *WithdrawHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseWithdraw(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewWithdrawHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolWithdraw]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &WithdrawHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
