package transaction

import (
	"github.com/IR-Digital-Token/x/chain/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type dataHandler struct {
	tx types.Transaction
	cb transactions.CallbackFn
}

func NewHandler(transaction types.Transaction, cb func(header types.Header, recipt *types.Receipt) error) transactions.Handler {
	newHandler := new(dataHandler)
	newHandler.tx = transaction
	newHandler.cb = cb
	return newHandler
}
func (t dataHandler) ID() common.Hash {
	return t.tx.Hash()
}

func (t dataHandler) HandleTransaction() func(header types.Header, recipt *types.Receipt) error {
	return func(header types.Header, recipt *types.Receipt) error {
		return t.cb(header, recipt)
	}
}
