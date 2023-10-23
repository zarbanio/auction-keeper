// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package osm

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type LogValueHandler struct {
	addr string
	binding  *Osm
	callback events.CallbackFn[OsmLogValue]
}

func (h *LogValueHandler) ID() string {
	return h.addr + ":" + "0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527"
}

func (h *LogValueHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLogValue(log.Log)
}

func (h *LogValueHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(OsmLogValue)
	if !ok {
		return errors.New("event type is not OsmLogValue")
	}
	return h.callback(log, e)
}

func (h *LogValueHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLogValue(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLogValueHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[OsmLogValue]) events.Handler {
	b, err := NewOsm(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LogValueHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
