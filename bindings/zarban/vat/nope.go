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

type NopeHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatNope]
}

func (h *NopeHandler) ID() string {
	return h.addr + ":" + "0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc"
}

func (h *NopeHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseNope(log.Log)
}

func (h *NopeHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatNope)
	if !ok {
		return errors.New("event type is not VatNope")
	}
	return h.callback(log, e)
}

func (h *NopeHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseNope(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewNopeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatNope]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NopeHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
