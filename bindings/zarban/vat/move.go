// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package vat

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type MoveHandler struct {
	addr string
	binding  *Vat
	callback events.CallbackFn[VatMove]
}

func (h *MoveHandler) ID() string {
	return h.addr + ":" + "0xdeb3a6837278f6e9914a507e4d73f08e841d8fca434fb97d4307b3b0d3d6b105"
}

func (h *MoveHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseMove(log.Log)
}

func (h *MoveHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(VatMove)
	if !ok {
		return errors.New("event type is not VatMove")
	}
	return h.callback(log, e)
}

func (h *MoveHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseMove(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewMoveHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[VatMove]) events.Handler {
	b, err := NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &MoveHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
