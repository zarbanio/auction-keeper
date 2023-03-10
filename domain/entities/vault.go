package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type Vault struct {
	UrnAddress common.Address `json:"urn_address"`
	Urn        Urn            `json:"urn"`
	IlkId      [32]byte       `json:"ilk_id"`
	IlkName    string         `json:"ilk_name"`
}

func IlkNameFromId(id [32]byte) string {
	return string(removeZeroBytes(id[:]))
}

func removeZeroBytes(arr []byte) []byte {
	var res []byte
	for _, elem := range arr {
		if elem != 0 {
			res = append(res, elem)
		}
	}
	return res
}
