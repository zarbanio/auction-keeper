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

type KissHandler struct {
	addr string
	binding  *Vow
	callback events.CallbackFn[VowKiss]
}

func (h *KissHandler) ID() string {
	return h.addr + ":" + "0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b"
}

func (h *KissHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseKiss(log.Log)
}

func (h *KissHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VowKiss)
	if !ok {
		return errors.New("event type is not VowKiss")
	}
	return h.callback(log, e)
}

func (h *KissHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseKiss(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewKissHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VowKiss]) events.Handler {
	b, err := NewVow(addr, eth)
	if err != nil {
		panic(err)
	}
	return &KissHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
