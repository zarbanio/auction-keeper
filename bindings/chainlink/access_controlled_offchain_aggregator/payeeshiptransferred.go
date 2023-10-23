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

type PayeeshipTransferredHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorPayeeshipTransferred]
}

func (h *PayeeshipTransferredHandler) ID() string {
	return h.addr + ":" + "0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3"
}

func (h *PayeeshipTransferredHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParsePayeeshipTransferred(log.Log)
}

func (h *PayeeshipTransferredHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorPayeeshipTransferred)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorPayeeshipTransferred")
	}
	return h.callback(log, e)
}

func (h *PayeeshipTransferredHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParsePayeeshipTransferred(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewPayeeshipTransferredHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorPayeeshipTransferred]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &PayeeshipTransferredHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
