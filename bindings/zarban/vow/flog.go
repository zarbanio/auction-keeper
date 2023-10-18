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

type FlogHandler struct {
	addr string
	binding  *Vow
	callback events.CallbackFn[VowFlog]
}

func (h *FlogHandler) ID() string {
	return h.addr + ":" + "0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753"
}

func (h *FlogHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlog(log.Log)
}

func (h *FlogHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VowFlog)
	if !ok {
		return errors.New("event type is not VowFlog")
	}
	return h.callback(log, e)
}

func (h *FlogHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlog(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFlogHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VowFlog]) events.Handler {
	b, err := NewVow(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlogHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
