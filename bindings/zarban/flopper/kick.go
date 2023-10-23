// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package flopper

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type KickHandler struct {
	addr string
	binding  *Flopper
	callback events.CallbackFn[FlopperKick]
}

func (h *KickHandler) ID() string {
	return h.addr + ":" + "0x0a90cdede34e48dec267bc60cbbd90898951b2fc42aa23620672bc93d73d8185"
}

func (h *KickHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseKick(log.Log)
}

func (h *KickHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(FlopperKick)
	if !ok {
		return errors.New("event type is not FlopperKick")
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

func NewKickHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[FlopperKick]) events.Handler {
	b, err := NewFlopper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &KickHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
