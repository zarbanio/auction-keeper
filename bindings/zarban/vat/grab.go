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

type GrabHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatGrab]
}

func (h *GrabHandler) ID() string {
	return h.addr + ":" + "0x1b2837fd40844c96cf39e52acaae7902fb74257fe20b1b7df5458b97d896c636"
}

func (h *GrabHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseGrab(log.Log)
}

func (h *GrabHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatGrab)
	if !ok {
		return errors.New("event type is not VatGrab")
	}
	return h.callback(log, e)
}

func (h *GrabHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseGrab(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewGrabHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatGrab]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &GrabHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
