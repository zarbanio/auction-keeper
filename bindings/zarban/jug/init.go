// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package jug

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type InitHandler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugInit]
}

func (h *InitHandler) ID() string {
	return h.addr + ":" + "0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555"
}

func (h *InitHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseInit(log.Log)
}

func (h *InitHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugInit)
	if !ok {
		return errors.New("event type is not JugInit")
	}
	return h.callback(log, e)
}

func (h *InitHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseInit(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewInitHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugInit]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &InitHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
