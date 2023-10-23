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

type IncreaseObservationCardinalityNextHandler struct {
	addr string
	binding  *Pool
	callback events.CallbackFn[PoolIncreaseObservationCardinalityNext]
}

func (h *IncreaseObservationCardinalityNextHandler) ID() string {
	return h.addr + ":" + "0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a"
}

func (h *IncreaseObservationCardinalityNextHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseIncreaseObservationCardinalityNext(log.Log)
}

func (h *IncreaseObservationCardinalityNextHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(PoolIncreaseObservationCardinalityNext)
	if !ok {
		return errors.New("event type is not PoolIncreaseObservationCardinalityNext")
	}
	return h.callback(log, e)
}

func (h *IncreaseObservationCardinalityNextHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseIncreaseObservationCardinalityNext(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewIncreaseObservationCardinalityNextHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[PoolIncreaseObservationCardinalityNext]) events.Handler {
	b, err := NewPool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &IncreaseObservationCardinalityNextHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
