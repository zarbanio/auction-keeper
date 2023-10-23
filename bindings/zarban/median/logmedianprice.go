// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package median

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type LogMedianPriceHandler struct {
	addr string
	binding  *Median
	callback events.CallbackFn[MedianLogMedianPrice]
}

func (h *LogMedianPriceHandler) ID() string {
	return h.addr + ":" + "0xb78ebc573f1f889ca9e1e0fb62c843c836f3d3a2e1f43ef62940e9b894f4ea4c"
}

func (h *LogMedianPriceHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLogMedianPrice(log.Log)
}

func (h *LogMedianPriceHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(MedianLogMedianPrice)
	if !ok {
		return errors.New("event type is not MedianLogMedianPrice")
	}
	return h.callback(log, e)
}

func (h *LogMedianPriceHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLogMedianPrice(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLogMedianPriceHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[MedianLogMedianPrice]) events.Handler {
	b, err := NewMedian(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LogMedianPriceHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
