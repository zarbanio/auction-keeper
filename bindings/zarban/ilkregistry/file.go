// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package ilkregistry

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type FileHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryFile]
}

func (h *FileHandler) ID() string {
	return h.addr + ":" + "0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba"
}

func (h *FileHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile(log.Log)
}

func (h *FileHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryFile)
	if !ok {
		return errors.New("event type is not IlkregistryFile")
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

func NewFileHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryFile]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
