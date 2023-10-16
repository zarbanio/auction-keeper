package loaders

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain/entities"
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
		IlkId:   ilkId,
		IlkName: entities.IlkNameFromId(ilkId),
	}
	return vault, nil

}
