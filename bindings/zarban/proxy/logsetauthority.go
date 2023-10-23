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

type LogSetAuthorityHandler struct {
	addr string
	binding  *Proxy
	callback events.CallbackFn[ProxyLogSetAuthority]
}

func (h *LogSetAuthorityHandler) ID() string {
	return h.addr + ":" + "0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4"
}

func (h *LogSetAuthorityHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLogSetAuthority(log.Log)
}

func (h *LogSetAuthorityHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ProxyLogSetAuthority)
	if !ok {
		return errors.New("event type is not ProxyLogSetAuthority")
	}
	return h.callback(log, e)
}

func (h *LogSetAuthorityHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLogSetAuthority(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLogSetAuthorityHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ProxyLogSetAuthority]) events.Handler {
	b, err := NewProxy(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LogSetAuthorityHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
