package account

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/ierc20"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

func balance(cfg configs.Config, secrets configs.Secrets) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}
	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}
	newSigner, err := signer.NewSigner(secrets.PrivateKey, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}
	sender := sender.NewSender(newSigner, postgresStore, eth)

	memCache := cache.NewMemCache()

	addressesLoader := loaders.NewAddressLoader(eth, memCache, cfg.Contracts.Deployment, cfg.Contracts.AddressProvider)
	addrs, err := addressesLoader.LoadAddresses(context.Background())
	if err != nil {
		log.Fatal("error loading addresses.", err)
	}

	ilksLoader := loaders.NewIlksLoader(
		eth,
		postgresStore,
		addrs["vat"],
		addrs["jug"],
		addrs["spot"],
		addrs["dog"],
		cfg.Contracts.IlkRegistry,
		cfg.Contracts.OsmRegistry,
	)

	ilks, err := ilksLoader.LoadIlks(context.Background())
	if err != nil {
		log.Fatal("error fetching ilks.", err)
	}

	vat, err := vat.NewVat(cfg.Contracts.Vat, eth)
	if err != nil {
		log.Fatal(err)
	}

	tokens := []domain.Token{
		{
			Decimals: 18,
			Symbol:   domain.ZAR,
			Address:  cfg.Contracts.ZAR,
		},
		{
			Decimals: 18,
			Symbol:   domain.WETH,
			Address:  cfg.Contracts.WETH,
		},
		{
			Decimals: 18,
			Symbol:   domain.WstETH,
			Address:  cfg.Contracts.WstETH,
		},
		{
			Decimals: 18,
			Symbol:   domain.DAI,
			Address:  cfg.Contracts.DAI,
		},
	}

	fmt.Println("Address:", sender.GetAddress().String())
	printSystemBalances(ilks, vat, sender.GetAddress())
	fmt.Println()
	printWalletBalances(eth, tokens, sender.GetAddress())
}

func printWalletBalances(eth *ethclient.Client, tokens []domain.Token, walletAddress common.Address) {
	fmt.Println("Wallet balances:")
	for _, token := range tokens {
		erc20, err := ierc20.NewIerc20(token.Address, eth)
		if err != nil {
			log.Fatal(err)
		}

		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: context.Background()}, walletAddress)
		if err != nil {
			log.Fatal(err)
		}

		poweredDecimal := math.BigPow(10, token.Decimals)
		balanceNormalized := decimal.NewFromBigInt(balance).Div(decimal.NewFromBigInt(poweredDecimal))
		fmt.Printf("%s: %s\n", token.Symbol, balanceNormalized.String())
	}

	nativeBalance, err := eth.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	wad := math.BigPow(10, 18)
	nativeBalanceNormalized := decimal.NewFromBigInt(nativeBalance).Div(decimal.NewFromBigInt(wad))
	fmt.Printf("%s: %s\n", "ETH", nativeBalanceNormalized.String())
}

func printSystemBalances(ilks []domain.Ilk, vat *vat.Vat, walletAddress common.Address) {
	wad := decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))
	zarInSystem, err := vat.Zar(&bind.CallOpts{Context: context.Background()}, walletAddress)
	if err != nil {
		log.Fatal(err)
	}
	zarBalance := new(big.Int).Div(zarInSystem, math.BigPow(10, 27))
	zarBalanceNormalized := decimal.NewFromBigInt(zarBalance).Div(wad)

	fmt.Println("System balances:")
	fmt.Printf("ZAR: %s\n", zarBalanceNormalized.String())

	for _, ilk := range ilks {
		balance, err := vat.Gem(&bind.CallOpts{Context: context.Background()}, ilk.Bytes32, walletAddress)
		if err != nil {
			log.Fatal(err)
		}
		balanceNormalized := decimal.NewFromBigInt(balance).Div(wad)
		fmt.Printf("%s(%s): %s\n", ilk.Symbol, ilk.Name, balanceNormalized.String())
	}
}
