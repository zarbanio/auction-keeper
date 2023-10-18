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

type HealHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatHeal]
}

func (h *HealHandler) ID() string {
	return h.addr + ":" + "0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7"
}

func (h *HealHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseHeal(log.Log)
}

func (h *HealHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatHeal)
	if !ok {
		return errors.New("event type is not VatHeal")
	}
	return h.callback(log, e)
}

func (h *HealHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseHeal(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewHealHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatHeal]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &HealHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
