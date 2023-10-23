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

type DenyHandler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugDeny]
}

func (h *DenyHandler) ID() string {
	return h.addr + ":" + "0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b"
}

func (h *DenyHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseDeny(log.Log)
}

func (h *DenyHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugDeny)
	if !ok {
		return errors.New("event type is not JugDeny")
	}
	return h.callback(log, e)
}

func (h *DenyHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseDeny(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewDenyHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugDeny]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &DenyHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
