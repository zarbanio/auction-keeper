// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package spot

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type PokeHandler struct {
	addr string
	binding  *Spot
	callback events.CallbackFn[SpotPoke]
}

func (h *PokeHandler) ID() string {
	return h.addr + ":" + "0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e"
}

func (h *PokeHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParsePoke(log.Log)
}

func (h *PokeHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(SpotPoke)
	if !ok {
		return errors.New("event type is not SpotPoke")
	}
	return h.callback(log, e)
}

func (h *PokeHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParsePoke(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewPokeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[SpotPoke]) events.Handler {
	b, err := NewSpot(addr, eth)
	if err != nil {
		panic(err)
	}
	return &PokeHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
