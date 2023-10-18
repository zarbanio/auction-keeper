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

type LinkTokenSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorLinkTokenSet]
}

func (h *LinkTokenSetHandler) ID() string {
	return h.addr + ":" + "0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a"
}

func (h *LinkTokenSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseLinkTokenSet(log.Log)
}

func (h *LinkTokenSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorLinkTokenSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorLinkTokenSet")
	}
	return h.callback(log, e)
}

func (h *LinkTokenSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseLinkTokenSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewLinkTokenSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorLinkTokenSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &LinkTokenSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
