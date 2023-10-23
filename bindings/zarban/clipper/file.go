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

type FileHandler struct {
	addr string
	binding  *Clipper
	callback events.CallbackFn[ClipperFile]
}

func (h *FileHandler) ID() string {
	return h.addr + ":" + "0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7"
}

func (h *FileHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile(log.Log)
}

func (h *FileHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ClipperFile)
	if !ok {
		return errors.New("event type is not ClipperFile")
	}
	return h.callback(log, e)
}

func (h *FileHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFileHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperFile]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
