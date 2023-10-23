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

type DepositHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolDeposit]
}

func (h *DepositHandler) ID() string {
	return h.addr + ":" + "0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951"
}

func (h *DepositHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseDeposit(log.Log)
}

func (h *DepositHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolDeposit)
	if !ok {
		return errors.New("event type is not LendingpoolDeposit")
	}
	return h.callback(log, e)
}

func (h *DepositHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseDeposit(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewDepositHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolDeposit]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &DepositHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
