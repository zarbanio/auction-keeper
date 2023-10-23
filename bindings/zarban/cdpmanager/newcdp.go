// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package cdpmanager

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type NewCdpHandler struct {
	addr string
	binding  *Cdpmanager
	callback events.CallbackFn[CdpmanagerNewCdp]
}

func (h *NewCdpHandler) ID() string {
	return h.addr + ":" + "0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2"
}

func (h *NewCdpHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseNewCdp(log.Log)
}

func (h *NewCdpHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(CdpmanagerNewCdp)
	if !ok {
		return errors.New("event type is not CdpmanagerNewCdp")
	}
	return h.callback(log, e)
}

func (h *NewCdpHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseNewCdp(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewNewCdpHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[CdpmanagerNewCdp]) events.Handler {
	b, err := NewCdpmanager(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NewCdpHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
