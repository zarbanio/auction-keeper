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

type ForkHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatFork]
}

func (h *ForkHandler) ID() string {
	return h.addr + ":" + "0x4b67161d2a4293b296b2f023c52ea4214353fa4f03e58973572faa00097dbd1e"
}

func (h *ForkHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFork(log.Log)
}

func (h *ForkHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatFork)
	if !ok {
		return errors.New("event type is not VatFork")
	}
	return h.callback(log, e)
}

func (h *ForkHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFork(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewForkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatFork]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ForkHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
