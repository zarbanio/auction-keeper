// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package proxy

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type LogSetOwnerHandler struct {
	addr string
	binding  *Proxy
	callback events.CallbackFn[ProxyLogSetOwner]
}

func (h *LogSetOwnerHandler) ID() string {
	return h.addr + ":" + "0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94"
}

func (h *LogSetOwnerHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLogSetOwner(log.Log)
}

func (h *LogSetOwnerHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ProxyLogSetOwner)
	if !ok {
		return errors.New("event type is not ProxyLogSetOwner")
	}
	return h.callback(log, e)
}

func (h *LogSetOwnerHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLogSetOwner(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLogSetOwnerHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ProxyLogSetOwner]) events.Handler {
	b, err := NewProxy(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LogSetOwnerHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
