// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package jug

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type FileHandler struct {
	addr string
	binding  *Jug
	callback events.CallbackFn[JugFile]
}

func (h *FileHandler) ID() string {
	return h.addr + ":" + "0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1"
}

func (h *FileHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile(log.Log)
}

func (h *FileHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(JugFile)
	if !ok {
		return errors.New("event type is not JugFile")
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

func NewFileHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[JugFile]) events.Handler {
	b, err := NewJug(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
