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

type NewTransmissionHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorNewTransmission]
}

func (h *NewTransmissionHandler) ID() string {
	return h.addr + ":" + "0xf6a97944f31ea060dfde0566e4167c1a1082551e64b60ecb14d599a9d023d451"
}

func (h *NewTransmissionHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseNewTransmission(log.Log)
}

func (h *NewTransmissionHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorNewTransmission)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorNewTransmission")
	}
	return h.callback(log, e)
}

func (h *NewTransmissionHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseNewTransmission(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewNewTransmissionHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorNewTransmission]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NewTransmissionHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
