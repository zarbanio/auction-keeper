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

type LiquidationCallHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolLiquidationCall]
}

func (h *LiquidationCallHandler) ID() string {
	return h.addr + ":" + "0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286"
}

func (h *LiquidationCallHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLiquidationCall(log.Log)
}

func (h *LiquidationCallHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolLiquidationCall)
	if !ok {
		return errors.New("event type is not LendingpoolLiquidationCall")
	}
	return h.callback(log, e)
}

func (h *LiquidationCallHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLiquidationCall(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLiquidationCallHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolLiquidationCall]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LiquidationCallHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
