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

type InitializeHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolInitialize]
}

func (h *InitializeHandler) ID() string {
	return h.addr + ":" + "0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95"
}

func (h *InitializeHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseInitialize(log.Log)
}

func (h *InitializeHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolInitialize)
	if !ok {
		return errors.New("event type is not PoolInitialize")
	}
	return h.callback(log, e)
}

func (h *InitializeHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseInitialize(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewInitializeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolInitialize]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &InitializeHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
