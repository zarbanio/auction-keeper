package loaders

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
)

type VowLoader struct {
	vow *vow.Vow
}

func NewVowLoader(eth *ethclient.Client, vowAddr common.Address) *VowLoader {

	v, err := vow.NewVow(vowAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &VowLoader{
		vow: v,
	}
}

func (vl *VowLoader) GetSinOf(ctx context.Context, era *big.Int) (*big.Int, error) {
	// sin is debt queue, mapping from era to sin;
	sin, err := vl.vow.Sin(&bind.CallOpts{Context: ctx}, era)
	if err != nil {
		return nil, err
	}

	return sin, nil
}

func (vl *VowLoader) GetSin(ctx context.Context) (*big.Int, error) {
	// Sin: Queued debt.
	sin, err := vl.vow.TotalSin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return sin, nil
}

func (vl *VowLoader) GetAsh(ctx context.Context) (*big.Int, error) {
	// Ash: On-auction debt.
	ash, err := vl.vow.Ash(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return ash, nil
}

func (vl *VowLoader) GetSump(ctx context.Context) (*big.Int, error) {
	// Sump: flop fixed lot size
	sump, err := vl.vow.Sump(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return sump, nil
}

func (vl *VowLoader) GetWait(ctx context.Context) (*big.Int, error) {
	// Wait: flop delay
	wait, err := vl.vow.Wait(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return wait, nil
}
