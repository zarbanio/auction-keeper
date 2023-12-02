package loaders

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/cdpmanager"
	"github.com/zarbanio/auction-keeper/bindings/zarban/getcdps"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type VaultLoader struct {
	store       store.IStore
	managerAddr common.Address
	getcdps     *getcdps.Getcdps
	manager     *cdpmanager.Cdpmanager
	vat         *vat.Vat
}

var (
	ErrIlkNotSupported = errors.New("ilk not supported")
)

type cdp struct {
	id  *big.Int
	urn common.Address
	ilk [32]byte
}

func NewVaultLoader(
	eth *ethclient.Client,
	store store.IStore,
	managerAddr,
	getcdpsAddr,
	vatAddr common.Address) *VaultLoader {

	m, err := cdpmanager.NewCdpmanager(managerAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	g, err := getcdps.NewGetcdps(getcdpsAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	v, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &VaultLoader{
		store:       store,
		vat:         v,
		getcdps:     g,
		manager:     m,
		managerAddr: managerAddr,
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

func (vl *VaultLoader) FetchVaultById(ctx context.Context, id *big.Int) (*domain.Vault, error) {
	ilk, err := vl.manager.Ilks(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}
	urn, err := vl.manager.Urns(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}
	c := cdp{
		id:  id,
		urn: urn,
		ilk: ilk,
	}

	return vl.fetchVaultByCDP(ctx, c)
}

func (vl *VaultLoader) fetchVaultByCDP(ctx context.Context, c cdp) (*domain.Vault, error) {
	urnInfo, err := vl.vat.Urns(&bind.CallOpts{Context: ctx}, c.ilk, c.urn)
	if err != nil {
		return nil, err
	}
	ilk, err := vl.store.GetIlkByName(ctx, domain.Bytes32ToString(c.ilk))
	if err != nil {
		if errors.Is(err, cache.ErrIlkNotFound) {
			return nil, ErrIlkNotSupported
		}
		return nil, err
	}

	proxy, err := vl.manager.Owns(&bind.CallOpts{Context: ctx}, c.id)
	if err != nil {
		return nil, err
	}

	collateral := math.Normalize(urnInfo.Ink, int64(math.WadDecimals))
	price := domain.NewCurrency(domain.ZAR, ilk.Price).SetValueInSymbol(domain.Toman, ilk.Price.Mul(decimal.NewFromInt(1)))
	collateralValue := collateral.Mul(price.ValueInSymbol(domain.ZAR))
	debt := debtValue(urnInfo.Art, ilk.Rate)

	vault := &domain.Vault{
		Id:                     c.id.Int64(),
		Owner:                  proxy,
		Urn:                    c.urn,
		LiquidationPrice:       liquidationPrice(collateral, debt, ilk.MinimumCollateralizationRatio).Mul(decimal.NewFromInt(1)),
		CollateralLocked:       collateral,
		CollateralizationRatio: collateralizationRatio(collateralValue, debt),
		Debt:                   debt,
		AvailableToWithdraw:    collateral.Sub(minSafeCollateralAmount(debt, ilk.MinimumCollateralizationRatio, price.Value())),
		AvailableToMint:        availableToMint(collateralValue, debt, ilk.MinimumCollateralizationRatio),
		IlkName:                ilk.Name,
	}
	return vault, nil
}

func (vl *VaultLoader) FetchAllVaults(ctx context.Context) ([]domain.Vault, error) {
	cdpi, err := vl.manager.Cdpi(&bind.CallOpts{Context: ctx})

	if err != nil {
		return []domain.Vault{}, err
	}

	var vaults []domain.Vault
	for i := int64(1); i <= int64(cdpi.Uint64()); i++ {
		vault, err := vl.FetchVaultById(ctx, big.NewInt(i))
		if err != nil {
			return []domain.Vault{}, err
		}
		vaults = append(vaults, *vault)
	}

	return vaults, nil
}
