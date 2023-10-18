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

type FoldHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFold]
}

func (h *FoldHandler) ID() string {
	return h.addr + ":" + "0x8e03d1ac49b6d4e90dd1c4e641ecc5ca76b7c07a487690b6897c0e5e374b19d2"
}

func (h *FoldHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFold(log.Log)
}

func (h *FoldHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFold)
	if !ok {
		return errors.New("event type is not VatFold")
	}
	return h.callback(log, e)
}

func (h *FoldHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFold(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFoldHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFold]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FoldHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
