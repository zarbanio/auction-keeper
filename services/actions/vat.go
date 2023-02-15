package actions

import "github.com/ethereum/go-ethereum/common"

func (s Actions) Hope(usr common.Address) (string, error) {

	opts, err := s.sender.GetOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vat.VatTransactor.Hope(opts, usr)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
