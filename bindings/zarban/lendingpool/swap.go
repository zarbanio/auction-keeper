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

type SwapHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolSwap]
}

func (h *SwapHandler) ID() string {
	return h.addr + ":" + "0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6"
}

func (h *SwapHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSwap(log.Log)
}

func (h *SwapHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolSwap)
	if !ok {
		return errors.New("event type is not LendingpoolSwap")
	}
	return h.callback(log, e)
}

func (h *SwapHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseSwap(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewSwapHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolSwap]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SwapHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
