// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type KickHandler struct {
	addr string
	binding  *Clipper
	callback events.CallbackFn[ClipperKick]
}

func (h *KickHandler) ID() string {
	return h.addr + ":" + "0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8"
}

func (h *KickHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseKick(log.Log)
}

func (h *KickHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ClipperKick)
	if !ok {
		return errors.New("event type is not ClipperKick")
	}
	return h.callback(log, e)
}

func (h *KickHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseKick(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewKickHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperKick]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &KickHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
