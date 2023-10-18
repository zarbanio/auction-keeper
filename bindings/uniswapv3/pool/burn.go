// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package pool

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type BurnHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolBurn]
}

func (h *BurnHandler) ID() string {
	return h.addr + ":" + "0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c"
}

func (h *BurnHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseBurn(log.Log)
}

func (h *BurnHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolBurn)
	if !ok {
		return errors.New("event type is not PoolBurn")
	}
	return h.callback(log, e)
}

func (h *BurnHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseBurn(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewBurnHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolBurn]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &BurnHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
