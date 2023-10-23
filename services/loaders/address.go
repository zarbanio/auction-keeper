package loaders

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/deployment"
	"github.com/zarbanio/auction-keeper/bindings/zarban/lendingpool_address_provider"
	"github.com/zarbanio/auction-keeper/cache"
)

type AddressLoader struct {
	store          cache.MemCache
	deployment     *deployment.Deployment
	deploymentAddr common.Address
	addrProvider   *lendingpool_address_provider.LendingpoolAddressProvider
}

func NewAddressLoader(
	eth *ethclient.Client,
	store cache.MemCache,
	deployAddr,
	addrProviderAddr common.Address,
) *AddressLoader {

	d, err := deployment.NewDeployment(deployAddr, eth)
	if err != nil {
		log.Fatal(err)
	}
	a, err := lendingpool_address_provider.NewLendingpoolAddressProvider(addrProviderAddr, eth)
	if err != nil {
		log.Fatal(err)
	}
	return &AddressLoader{
		deployment:     d,
		deploymentAddr: deployAddr,
		addrProvider:   a,
		store:          store,
	}
}

func (a *AddressLoader) LoadAddresses(ctx context.Context) (map[string]common.Address, error) {
	lendingPoolAddr, err := a.addrProvider.GetLendingPool(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting lending pool. %w", err)
	}
	emergencyAdmin, err := a.addrProvider.GetEmergencyAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting emergency admin. %w", err)
	}
	collateralManager, err := a.addrProvider.GetLendingPoolCollateralManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting lending pool collateral manager. %w", err)
	}
	configurator, err := a.addrProvider.GetLendingPoolConfigurator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting lending pool configurator")
	}
	admin, err := a.addrProvider.GetPoolAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting pool admin. %w", err)
	}
	oracle, err := a.addrProvider.GetPriceOracle(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting price oracle. %w", err)
	}
	vat, err := a.deployment.Vat(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting vat. %w", err)
	}
	jug, err := a.deployment.Jug(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting jug. %w", err)
	}
	spot, err := a.deployment.Spotter(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting spotter. %w", err)
	}
	dog, err := a.deployment.Dog(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting dog. %w", err)
	}
	end, err := a.deployment.End(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting end. %w", err)
	}
	vow, err := a.deployment.Vow(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting vow. %w", err)
	}

	zar, err := a.deployment.Zar(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting sim. %w", err)
	}
	zarJoin, err := a.deployment.ZarJoin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting sim join. %w", err)
	}

	return map[string]common.Address{
		"deployment":         a.deploymentAddr,
		"lending_pool":       lendingPoolAddr,
		"emergency_admin":    emergencyAdmin,
		"collateral_manager": collateralManager,
		"configurator":       configurator,
		"admin":              admin,
		"price_oracle":       oracle,
		"vat":                vat,
		"jug":                jug,
		"spot":               spot,
		"dog":                dog,
		"end":                end,
		"vow":                vow,
		"zar":                zar,
		"zar_join":           zarJoin,
	}, nil
}

func (a *AddressLoader) SaveAddresses(ctx context.Context, addresses map[string]common.Address) error {
	return a.store.SaveAddresses(ctx, addresses)
}
