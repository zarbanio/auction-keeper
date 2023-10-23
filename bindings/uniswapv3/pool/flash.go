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

type FlashHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolFlash]
}

func (h *FlashHandler) ID() string {
	return h.addr + ":" + "0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633"
}

func (h *FlashHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlash(log.Log)
}

func (h *FlashHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolFlash)
	if !ok {
		return errors.New("event type is not PoolFlash")
	}
	return h.callback(log, e)
}

func (h *FlashHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlash(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFlashHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolFlash]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlashHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
