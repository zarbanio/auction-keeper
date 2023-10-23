// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package vow

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type FessHandler struct {
	addr string
	binding  *Vow
	callback events.CallbackFn[VowFess]
}

func (h *FessHandler) ID() string {
	return h.addr + ":" + "0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba"
}

func (h *FessHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFess(log.Log)
}

func (h *FessHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VowFess)
	if !ok {
		return errors.New("event type is not VowFess")
	}
	return h.callback(log, e)
}

func (h *FessHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFess(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFessHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VowFess]) events.Handler {
	b, err := NewVow(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FessHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
