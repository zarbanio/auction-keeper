package transaction

import (
	"math/big"
)

func (s Sender) Flop() (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vow.VowTransactor.Flop(opts)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s Sender) Flog(era *big.Int) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vow.VowTransactor.Flog(opts, era)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s Sender) Kiss(rad *big.Int) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vow.VowTransactor.Kiss(opts, rad)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s Sender) Heal(rad *big.Int) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.vow.VowTransactor.Heal(opts, rad)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
