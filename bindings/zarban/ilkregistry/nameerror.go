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

type NameErrorHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryNameError]
}

func (h *NameErrorHandler) ID() string {
	return h.addr + ":" + "0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a"
}

func (h *NameErrorHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseNameError(log.Log)
}

func (h *NameErrorHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryNameError)
	if !ok {
		return errors.New("event type is not IlkregistryNameError")
	}
	return h.callback(log, e)
}

func (h *NameErrorHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseNameError(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewNameErrorHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryNameError]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NameErrorHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
