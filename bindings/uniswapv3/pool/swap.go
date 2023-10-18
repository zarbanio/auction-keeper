// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package pool

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type SwapHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolSwap]
}

func (h *SwapHandler) ID() string {
	return h.addr + ":" + "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
}

func (h *SwapHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSwap(log.Log)
}

func (h *SwapHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolSwap)
	if !ok {
		return errors.New("event type is not PoolSwap")
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

func NewSwapHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolSwap]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SwapHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
