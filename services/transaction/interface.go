package transaction

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ISender interface {
	GetOpts() (*bind.TransactOpts, error)
	GetAddress() common.Address
}
