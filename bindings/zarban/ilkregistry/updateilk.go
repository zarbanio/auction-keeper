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

type UpdateIlkHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryUpdateIlk]
}

func (h *UpdateIlkHandler) ID() string {
	return h.addr + ":" + "0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa"
}

func (h *UpdateIlkHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseUpdateIlk(log.Log)
}

func (h *UpdateIlkHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryUpdateIlk)
	if !ok {
		return errors.New("event type is not IlkregistryUpdateIlk")
	}
	return h.callback(log, e)
}

func (h *UpdateIlkHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseUpdateIlk(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewUpdateIlkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryUpdateIlk]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &UpdateIlkHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
