package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type Vault struct {
	UrnAddress common.Address `json:"urn_address"`
	Urn        Urn            `json:"urn"`
	Ilk        Ilk            `json:"ilk"`
}
