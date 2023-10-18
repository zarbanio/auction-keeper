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

type SuckHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatSuck]
}

func (h *SuckHandler) ID() string {
	return h.addr + ":" + "0x02d16dda43fd89f02e33ce23ecf0251cdc426807cc72ae74d37e8d3681dae7e5"
}

func (h *SuckHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSuck(log.Log)
}

func (h *SuckHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatSuck)
	if !ok {
		return errors.New("event type is not VatSuck")
	}
	return h.callback(log, e)
}

func (h *SuckHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseSuck(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewSuckHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatSuck]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SuckHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
