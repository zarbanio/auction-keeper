package transaction

import (
	"github.com/IR-Digital-Token/x/chain/transactions"
)

func (r Sender) WatchTransactionHash(tx transactions.Handler) {
	r.indexer.WatchTx(tx)
}

func (r Sender) UnWatchWatchTransactionHash(tx transactions.Handler) {
	r.indexer.UnWatchTx(tx)
}
