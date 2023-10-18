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

type FluxHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFlux]
}

func (h *FluxHandler) ID() string {
	return h.addr + ":" + "0x5718eae79ffb8b6c98c497e5029a903705cf6a33a17aaab32de7fe198d8d8a0d"
}

func (h *FluxHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlux(log.Log)
}

func (h *FluxHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFlux)
	if !ok {
		return errors.New("event type is not VatFlux")
	}
	return h.callback(log, e)
}

func (h *FluxHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlux(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFluxHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFlux]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FluxHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
