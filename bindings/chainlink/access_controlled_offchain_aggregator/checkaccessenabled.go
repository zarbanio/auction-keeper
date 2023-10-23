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

type CheckAccessEnabledHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorCheckAccessEnabled]
}

func (h *CheckAccessEnabledHandler) ID() string {
	return h.addr + ":" + "0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480"
}

func (h *CheckAccessEnabledHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseCheckAccessEnabled(log.Log)
}

func (h *CheckAccessEnabledHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorCheckAccessEnabled)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorCheckAccessEnabled")
	}
	return h.callback(log, e)
}

func (h *CheckAccessEnabledHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseCheckAccessEnabled(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewCheckAccessEnabledHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorCheckAccessEnabled]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CheckAccessEnabledHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
