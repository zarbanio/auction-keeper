package loaders

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type VaultLoader struct {
	vat *vat.Vat
}

func NewVaultLoader(eth *ethclient.Client, vatAddr common.Address) *VaultLoader {

	v, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &VaultLoader{
		vat: v,
	}
}

func (vl *VaultLoader) GetIlkById(ctx context.Context, ilkId [32]byte) (*entities.VatIlk, error) {
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

func (vl *VaultLoader) GetVaultByIlkUrn(ctx context.Context, ilkId [32]byte, urn common.Address) (*entities.Vault, error) {
	urnInfo, err := vl.vat.Urns(&bind.CallOpts{Context: ctx}, ilkId, urn)
	if err != nil {
		return nil, err
	}

	vault := &entities.Vault{
		UrnAddress: urn,
		Urn: entities.Urn{
			Ink: urnInfo.Ink,
			Art: urnInfo.Art,
		},
		IlkId: ilkId,
	}
	return vault, nil

}
