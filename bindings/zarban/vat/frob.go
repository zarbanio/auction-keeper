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

type FrobHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFrob]
}

func (h *FrobHandler) ID() string {
	return h.addr + ":" + "0xe37707842c8387f7c3c357f1d6c5bf57084e681573bdda024fae70cf0ecde80e"
}

func (h *FrobHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFrob(log.Log)
}

func (h *FrobHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFrob)
	if !ok {
		return errors.New("event type is not VatFrob")
	}
	return h.callback(log, e)
}

func (h *FrobHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFrob(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFrobHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFrob]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FrobHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
