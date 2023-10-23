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

type HopeHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatHope]
}

func (h *HopeHandler) ID() string {
	return h.addr + ":" + "0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43"
}

func (h *HopeHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseHope(log.Log)
}

func (h *HopeHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatHope)
	if !ok {
		return errors.New("event type is not VatHope")
	}
	return h.callback(log, e)
}

func (h *HopeHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseHope(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewHopeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatHope]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &HopeHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
