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

type CollectProtocolHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolCollectProtocol]
}

func (h *CollectProtocolHandler) ID() string {
	return h.addr + ":" + "0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151"
}

func (h *CollectProtocolHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseCollectProtocol(log.Log)
}

func (h *CollectProtocolHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolCollectProtocol)
	if !ok {
		return errors.New("event type is not PoolCollectProtocol")
	}
	return h.callback(log, e)
}

func (h *CollectProtocolHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseCollectProtocol(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewCollectProtocolHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolCollectProtocol]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CollectProtocolHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
