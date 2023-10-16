package loaders

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/flopper"
)

type FlopperLoader struct {
	flopper *flopper.Flopper
}

func NewFlopperLoader(eth *ethclient.Client, flopperAddr common.Address) *FlopperLoader {

	v, err := flopper.NewFlopper(flopperAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &FlopperLoader{
		flopper: v,
	}
}

func (fl *FlopperLoader) IsRely(ctx context.Context, address common.Address) (bool, error) {
	// sin: unbacked stablecoin (system debt, not belonging to any urn).
	ward, err := fl.flopper.Wards(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return false, err
	}

	// if ward[address] == 1 return true else return false
	return ward.Cmp(big.NewInt(1)) == 0, nil
}
