package vaults

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/store"
)

func listVaults(cfg configs.Config, secrets configs.Secrets) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)

	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}

	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	memCache := cache.NewMemCache()
	addressesLoader := loaders.NewAddressLoader(eth, memCache, cfg.Contracts.Deployment, cfg.Contracts.AddressProvider)
	addrs, err := addressesLoader.LoadAddresses(context.Background())
	if err != nil {
		log.Fatal("error loading addresses.", err)
	}

	addrs["cdp_manager"] = cfg.Contracts.CDPManager
	addrs["ilk_registry"] = cfg.Contracts.IlkRegistry
	addrs["eth_a_join"] = cfg.Contracts.ETHAJoin
	addrs["eth_b_join"] = cfg.Contracts.ETHBJoin
	addrs["dai_a_join"] = cfg.Contracts.DAIAJoin
	addrs["dai_b_join"] = cfg.Contracts.DAIBJoin
	addrs["wsteth_a_join"] = cfg.Contracts.WstETHAJoin
	addrs["dai_median"] = cfg.Contracts.DAIMedian
	addrs["eth_median"] = cfg.Contracts.ETHMedian
	addrs["wsteth_median"] = cfg.Contracts.WstETHMedian

	ilksLoader := loaders.NewIlksLoader(
		eth,
		postgresStore,
		addrs["vat"],
		addrs["jug"],
		addrs["spot"],
		addrs["dog"],
		addrs["ilk_registry"],
		[]common.Address{
			addrs["eth_a_join"],
			addrs["eth_b_join"],
			addrs["dai_a_join"],
			addrs["dai_b_join"],
			addrs["wsteth_a_join"],
		},
		map[common.Address]common.Address{
			cfg.Contracts.DAI:    addrs["dai_median"],
			cfg.Contracts.WETH:   addrs["eth_median"],
			cfg.Contracts.WstETH: addrs["wsteth_median"],
		},
	)

	ilks, err := ilksLoader.LoadIlks(context.Background())
	if err != nil {
		log.Fatal("error loading ilks.", err)
	}

	for _, ilk := range ilks {
		err = postgresStore.CreateOrUpdateIlk(context.Background(), ilk)
		if err != nil {
			log.Fatal("error saving ilks.\n", err)
		}
	}

	vaultLoader := loaders.NewVaultLoader(
		eth,
		postgresStore,
		addrs["cdp_manager"],
		addrs["vat"],
	)

	vaults, err := vaultLoader.FetchAllVaults(context.Background())

	if err != nil {
		log.Fatal("error while fetching all vaults. ", err)
	}

	fmt.Println("--------------------------------------------------")
	for _, v := range vaults {
		fmt.Println("Id:                       ", v.Id)
		fmt.Println("Owner:                    ", v.Owner.String())
		fmt.Println("Urn:                      ", v.Urn.String())
		fmt.Println("LiquidationPrice:         ", v.LiquidationPrice.String())
		fmt.Println("CollateralLocked:         ", v.CollateralLocked.String())
		fmt.Println("CollateralizationRatio:   ", v.CollateralizationRatio.String())
		fmt.Println("Debt:                     ", v.Debt.String())
		fmt.Println("AvailableToWithdraw:      ", v.AvailableToWithdraw.String())
		fmt.Println("AvailableToMint:          ", v.AvailableToMint.String())
		fmt.Println("IlkName:                  ", v.IlkName)
		fmt.Println("--------------------------------------------------")

	}
}

func listVault(cfg configs.Config, secrets configs.Secrets, vaultId *big.Int) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)

	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}

	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	memCache := cache.NewMemCache()
	addressesLoader := loaders.NewAddressLoader(eth, memCache, cfg.Contracts.Deployment, cfg.Contracts.AddressProvider)
	addrs, err := addressesLoader.LoadAddresses(context.Background())
	if err != nil {
		log.Fatal("error loading addresses.", err)
	}

	addrs["cdp_manager"] = cfg.Contracts.CDPManager
	addrs["ilk_registry"] = cfg.Contracts.IlkRegistry
	addrs["eth_a_join"] = cfg.Contracts.ETHAJoin
	addrs["eth_b_join"] = cfg.Contracts.ETHBJoin
	addrs["dai_a_join"] = cfg.Contracts.DAIAJoin
	addrs["dai_b_join"] = cfg.Contracts.DAIBJoin
	addrs["wsteth_a_join"] = cfg.Contracts.WstETHAJoin
	addrs["dai_median"] = cfg.Contracts.DAIMedian
	addrs["eth_median"] = cfg.Contracts.ETHMedian
	addrs["wsteth_median"] = cfg.Contracts.WstETHMedian

	ilksLoader := loaders.NewIlksLoader(
		eth,
		postgresStore,
		addrs["vat"],
		addrs["jug"],
		addrs["spot"],
		addrs["dog"],
		addrs["ilk_registry"],
		[]common.Address{
			addrs["eth_a_join"],
			addrs["eth_b_join"],
			addrs["dai_a_join"],
			addrs["dai_b_join"],
			addrs["wsteth_a_join"],
		},
		map[common.Address]common.Address{
			cfg.Contracts.DAI:    addrs["dai_median"],
			cfg.Contracts.WETH:   addrs["eth_median"],
			cfg.Contracts.WstETH: addrs["wsteth_median"],
		},
	)

	ilks, err := ilksLoader.LoadIlks(context.Background())
	if err != nil {
		log.Fatal("error loading ilks.", err)
	}

	for _, ilk := range ilks {
		err = postgresStore.CreateOrUpdateIlk(context.Background(), ilk)
		if err != nil {
			log.Fatal("error saving ilks.\n", err)
		}
	}

	vaultLoader := loaders.NewVaultLoader(
		eth,
		postgresStore,
		addrs["cdp_manager"],
		addrs["vat"],
	)

	vault, err := vaultLoader.FetchVaultById(context.Background(), vaultId)

	if err != nil {
		log.Fatal("error while fetching the vault. ", err)
	}

	fmt.Println("Id:                       ", vault.Id)
	fmt.Println("Owner:                    ", vault.Owner.String())
	fmt.Println("Urn:                      ", vault.Urn.String())
	fmt.Println("LiquidationPrice:         ", vault.LiquidationPrice.String())
	fmt.Println("CollateralLocked:         ", vault.CollateralLocked.String())
	fmt.Println("CollateralizationRatio:   ", vault.CollateralizationRatio.String())
	fmt.Println("Debt:                     ", vault.Debt.String())
	fmt.Println("AvailableToWithdraw:      ", vault.AvailableToWithdraw.String())
	fmt.Println("AvailableToMint:          ", vault.AvailableToMint.String())
	fmt.Println("IlkName:                  ", vault.IlkName)

}
