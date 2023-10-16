package store

import "errors"

var (
	ErrLogNotFound                  = errors.New("log not found")
	ErrTransactionNotFound          = errors.New("transaction not found")
	ErrMedianLogPriceMedianNotFound = errors.New("median logPriceMedain not found")
)
