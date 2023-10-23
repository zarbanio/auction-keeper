// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package vat

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type File0Handler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFile0]
}

func (h *File0Handler) ID() string {
	return h.addr + ":" + "0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1"
}

func (h *File0Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile0(log.Log)
}

func (h *File0Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFile0)
	if !ok {
		return errors.New("event type is not VatFile0")
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

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFile0]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
