package transaction

import (
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/ethereum/go-ethereum/common"
)

func (s Sender) Hope(vat *vat.Vat, usr common.Address) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := vat.VatTransactor.Hope(opts, usr)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
