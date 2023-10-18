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

type PayeeshipTransferRequestedHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorPayeeshipTransferRequested]
}

func (h *PayeeshipTransferRequestedHandler) ID() string {
	return h.addr + ":" + "0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367"
}

func (h *PayeeshipTransferRequestedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParsePayeeshipTransferRequested(log.Log)
}

func (h *PayeeshipTransferRequestedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorPayeeshipTransferRequested)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorPayeeshipTransferRequested")
	}
	return h.callback(log, e)
}

func (h *PayeeshipTransferRequestedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParsePayeeshipTransferRequested(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewPayeeshipTransferRequestedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorPayeeshipTransferRequested]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &PayeeshipTransferRequestedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
