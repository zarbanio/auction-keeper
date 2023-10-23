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

type DripHandler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugDrip]
}

func (h *DripHandler) ID() string {
	return h.addr + ":" + "0x2fbd8559330c6faf38d4ee435b8463427bd58cd6bc7edd1a4e6f3ab6b57de71e"
}

func (h *DripHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseDrip(log.Log)
}

func (h *DripHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugDrip)
	if !ok {
		return errors.New("event type is not JugDrip")
	}
	return h.callback(log, e)
}

func (h *DripHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseDrip(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewDripHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugDrip]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &DripHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
