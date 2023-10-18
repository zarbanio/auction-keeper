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

type OraclePaidHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorOraclePaid]
}

func (h *OraclePaidHandler) ID() string {
	return h.addr + ":" + "0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c"
}

func (h *OraclePaidHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseOraclePaid(log.Log)
}

func (h *OraclePaidHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorOraclePaid)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorOraclePaid")
	}
	return h.callback(log, e)
}

func (h *OraclePaidHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseOraclePaid(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewOraclePaidHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorOraclePaid]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &OraclePaidHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
