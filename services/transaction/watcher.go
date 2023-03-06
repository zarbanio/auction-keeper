package transaction

import (
	"github.com/IR-Digital-Token/x/chain/transactions"
)

func (s Sender) WatchTransactionHash(tx transactions.Handler) {
	s.indexer.WatchTx(tx)
}

func (s Sender) UnWatchWatchTransactionHash(tx transactions.Handler) {
	s.indexer.UnWatchTx(tx)
}
