// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package pool

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type SetFeeProtocolHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolSetFeeProtocol]
}

func (h *SetFeeProtocolHandler) ID() string {
	return h.addr + ":" + "0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133"
}

func (h *SetFeeProtocolHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSetFeeProtocol(log.Log)
}

func (h *SetFeeProtocolHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolSetFeeProtocol)
	if !ok {
		return errors.New("event type is not PoolSetFeeProtocol")
	}
	return h.callback(log, e)
}

func (h *SetFeeProtocolHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseSetFeeProtocol(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewSetFeeProtocolHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolSetFeeProtocol]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SetFeeProtocolHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
