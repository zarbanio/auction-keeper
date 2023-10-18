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

type YankHandler struct {
	addr string
	binding  *Clipper
	callback events.CallbackFn[ClipperYank]
}

func (h *YankHandler) ID() string {
	return h.addr + ":" + "0x2c5d2826eb5903b8fc201cf48094b858f42f61c7eaac9aaf43ebed490138144e"
}

func (h *YankHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseYank(log.Log)
}

func (h *YankHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ClipperYank)
	if !ok {
		return errors.New("event type is not ClipperYank")
	}
	return h.callback(log, e)
}

func (h *YankHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseYank(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewYankHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperYank]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &YankHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
