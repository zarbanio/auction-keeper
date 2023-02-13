package transaction

import "github.com/ethereum/go-ethereum/common"

func (s Sender) Bark(ilk [32]byte, urn common.Address) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := s.dog.DogTransactor.Bark(opts, ilk, urn, s.GetAddress())
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
