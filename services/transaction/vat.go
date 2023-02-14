package transaction

import "github.com/ethereum/go-ethereum/common"

func (s Sender) Hope(usr common.Address) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vat.VatTransactor.Hope(opts, usr)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
