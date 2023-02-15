package transaction

import (
	"github.com/IR-Digital-Token/x/chain/transactions"
)

func (r Sender) watchTransactionHash(tx transactions.Handler) {
	r.indexer.WatchTx(tx)
}

func (r Sender) unWatchWatchTransactionHash(tx transactions.Handler) {
	r.indexer.UnWatchTx(tx)
}
