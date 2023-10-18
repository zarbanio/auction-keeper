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

type MintHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolMint]
}

func (h *MintHandler) ID() string {
	return h.addr + ":" + "0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde"
}

func (h *MintHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseMint(log.Log)
}

func (h *MintHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolMint)
	if !ok {
		return errors.New("event type is not PoolMint")
	}
	return h.callback(log, e)
}

func (h *MintHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseMint(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewMintHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolMint]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &MintHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
