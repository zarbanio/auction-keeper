// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package access_controlled_offchain_aggregator

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type RequesterAccessControllerSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorRequesterAccessControllerSet]
}

func (h *RequesterAccessControllerSetHandler) ID() string {
	return h.addr + ":" + "0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634"
}

func (h *RequesterAccessControllerSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRequesterAccessControllerSet(log.Log)
}

func (h *RequesterAccessControllerSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorRequesterAccessControllerSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorRequesterAccessControllerSet")
	}
	return h.callback(log, e)
}

func (h *RequesterAccessControllerSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRequesterAccessControllerSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRequesterAccessControllerSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorRequesterAccessControllerSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RequesterAccessControllerSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
