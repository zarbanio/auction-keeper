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

type CollectHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolCollect]
}

func (h *CollectHandler) ID() string {
	return h.addr + ":" + "0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0"
}

func (h *CollectHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseCollect(log.Log)
}

func (h *CollectHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolCollect)
	if !ok {
		return errors.New("event type is not PoolCollect")
	}
	return h.callback(log, e)
}

func (h *CollectHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseCollect(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewCollectHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolCollect]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CollectHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
