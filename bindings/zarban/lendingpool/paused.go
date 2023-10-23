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

type PausedHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolPaused]
}

func (h *PausedHandler) ID() string {
	return h.addr + ":" + "0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752"
}

func (h *PausedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParsePaused(log.Log)
}

func (h *PausedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolPaused)
	if !ok {
		return errors.New("event type is not LendingpoolPaused")
	}
	return h.callback(log, e)
}

func (h *PausedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParsePaused(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewPausedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolPaused]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &PausedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
