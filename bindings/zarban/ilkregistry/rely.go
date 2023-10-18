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

type RelyHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryRely]
}

func (h *RelyHandler) ID() string {
	return h.addr + ":" + "0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60"
}

func (h *RelyHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRely(log.Log)
}

func (h *RelyHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryRely)
	if !ok {
		return errors.New("event type is not IlkregistryRely")
	}
	return h.callback(log, e)
}

func (h *RelyHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRely(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRelyHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryRely]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RelyHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
