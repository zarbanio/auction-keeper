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

type File0Handler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugFile0]
}

func (h *File0Handler) ID() string {
	return h.addr + ":" + "0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7"
}

func (h *File0Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile0(log.Log)
}

func (h *File0Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugFile0)
	if !ok {
		return errors.New("event type is not JugFile0")
	}
	return h.callback(log, e)
}

func (h *File0Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile0(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugFile0]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
