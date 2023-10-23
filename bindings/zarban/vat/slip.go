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

type SlipHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatSlip]
}

func (h *SlipHandler) ID() string {
	return h.addr + ":" + "0x0d5f62756a04d37a9bb68fd20b97c7c6a03e96ab87385a99f99c2463157dba4e"
}

func (h *SlipHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSlip(log.Log)
}

func (h *SlipHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatSlip)
	if !ok {
		return errors.New("event type is not VatSlip")
	}
	return h.callback(log, e)
}

func (h *SlipHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseSlip(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewSlipHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatSlip]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SlipHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
