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
	cache          cache.ICache
	deployment     *deployment.Deployment
	deploymentAddr common.Address
	addrProvider   *lendingpool_address_provider.LendingpoolAddressProvider
}

func NewAddressLoader(
	eth *ethclient.Client,
	store cache.ICache,
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
		cache:          store,
	}
}

func (a *AddressLoader) LoadAddresses(ctx context.Context) (map[string]common.Address, error) {
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
		"deployment": a.deploymentAddr,
		"vat":        vat,
		"jug":        jug,
		"spot":       spot,
		"dog":        dog,
		"end":        end,
		"vow":        vow,
		"zar":        zar,
		"zar_join":   zarJoin,
	}, nil
}

func (a *AddressLoader) SaveAddresses(ctx context.Context, addresses map[string]common.Address) error {
	return a.cache.SaveAddresses(ctx, addresses)
}
