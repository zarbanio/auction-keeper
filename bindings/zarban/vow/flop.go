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

type FlopHandler struct {
	addr string
	binding  *Vow
	callback events.CallbackFn[VowFlop]
}

func (h *FlopHandler) ID() string {
	return h.addr + ":" + "0x2a6fd69fb4120226e0d810a9039fc9a40879738ce61c80379d760f4614f6680d"
}

func (h *FlopHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFlop(log.Log)
}

func (h *FlopHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VowFlop)
	if !ok {
		return errors.New("event type is not VowFlop")
	}
	return h.callback(log, e)
}

func (h *FlopHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFlop(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFlopHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VowFlop]) events.Handler {
	b, err := NewVow(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlopHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
