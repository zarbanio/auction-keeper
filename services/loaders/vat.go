package loaders

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/vat"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type VatLoader struct {
	vat *vat.Vat
}

func NewVatLoader(eth *ethclient.Client, vatAddr common.Address) *VatLoader {

	v, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &VatLoader{
		vat: v,
	}
}

func (vl *VatLoader) GetIlkById(ctx context.Context, ilkId [32]byte) (*entities.VatIlk, error) {
	ilkVatInfo, err := vl.vat.Ilks(&bind.CallOpts{Context: ctx}, ilkId)
	if err != nil {
		return nil, err
	}

	ilkName := string(ilkId[:])
	ilk := &entities.VatIlk{
		Id:   ilkId,
		Name: ilkName,
		Art:  ilkVatInfo.Art,
		Rate: ilkVatInfo.Rate,
		Spot: ilkVatInfo.Spot,
		Line: ilkVatInfo.Line,
		Dust: ilkVatInfo.Dust,
	}

	return ilk, nil
}

func (vl *VatLoader) GetZarBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := vl.vat.Zar(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (vl *VatLoader) GetSin(ctx context.Context, address common.Address) (*big.Int, error) {
	// sin: unbacked stablecoin (system debt, not belonging to any urn).
	sin, err := vl.vat.Sin(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}

	return sin, nil
}
