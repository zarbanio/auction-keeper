// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package zarjoin

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type JoinHandler struct {
	addr string
	binding  *Zarjoin
	callback events.CallbackFn[ZarjoinJoin]
}

func (h *JoinHandler) ID() string {
	return h.addr + ":" + "0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95"
}

func (h *JoinHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseJoin(log.Log)
}

func (h *JoinHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ZarjoinJoin)
	if !ok {
		return errors.New("event type is not ZarjoinJoin")
	}
	return h.callback(log, e)
}

func (h *JoinHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseJoin(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewJoinHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ZarjoinJoin]) events.Handler {
	b, err := NewZarjoin(addr, eth)
	if err != nil {
		panic(err)
	}
	return &JoinHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
