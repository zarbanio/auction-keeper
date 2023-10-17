package domain

import "github.com/ethereum/go-ethereum/common"

type Proxy struct {
	Address   common.Address `json:"address"`
	Owner     common.Address `json:"owner"`
	Authority common.Address `json:"authority"`
}
