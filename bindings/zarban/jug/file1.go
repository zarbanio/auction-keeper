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

type File1Handler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugFile1]
}

func (h *File1Handler) ID() string {
	return h.addr + ":" + "0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba"
}

func (h *File1Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile1(log.Log)
}

func (h *File1Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugFile1)
	if !ok {
		return errors.New("event type is not JugFile1")
	}
	return h.callback(log, e)
}

func (h *File1Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile1(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile1Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugFile1]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File1Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
