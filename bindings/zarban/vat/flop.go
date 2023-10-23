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

type FlopHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFlop]
}

func (h *FlopHandler) ID() string {
	return h.addr + ":" + "0xfa3d951cbf852d2a9cc2dfc9fc6b57914afbf57597ecf432c403ed2d74124b2c"
}

func (h *FlopHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlop(log.Log)
}

func (h *FlopHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFlop)
	if !ok {
		return errors.New("event type is not VatFlop")
	}
	return h.callback(log, e)
}

func (h *FlopHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlop(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFlopHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFlop]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlopHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
