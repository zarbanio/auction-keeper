package cmd

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/bark"
	"github.com/zarbanio/auction-keeper/services/cachedeth"
	"github.com/zarbanio/auction-keeper/services/loaders"
	loggerPkg "github.com/zarbanio/auction-keeper/services/logger"
	"github.com/zarbanio/auction-keeper/services/redo"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/services/take"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
	"github.com/zarbanio/auction-keeper/store"
)

func Execute() {
	cfg := configs.ReadConfig("config.goerli.yaml")
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	logger := loggerPkg.NewLogger(context.Background(), postgresStore)

	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		log.Fatal(err)
	}

	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	bcache := cachedeth.NewBlockCache(eth)
	memCache := cache.NewMemCache()
	ceth := cachedeth.NewEthProxy(eth, postgresStore, bcache)

	newSigner, err := signer.NewSigner(cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}
	sender := sender.NewSender(newSigner, postgresStore, eth)

	addressesLoader := loaders.NewAddressLoader(eth, memCache, cfg.Contracts.Deployment, cfg.Contracts.AddressProvider)
	addrs, err := addressesLoader.LoadAddresses(context.Background())
	if err != nil {
		log.Fatal("error loading addresses.", err)
	}

	quoter, err := uniswap_v3.NewUniswapV3Quoter(eth, cfg.Contracts.UniswapV3Quoter)
	if err != nil {
		log.Fatal("error loading uniswap v3 quoter.", err)
	}

	addrs["cdp_manager"] = cfg.Contracts.CDPManager
	addrs["get_cdps"] = cfg.Contracts.GetCDPs
	addrs["ilk_registry"] = cfg.Contracts.IlkRegistry
	addrs["eth_a_join"] = cfg.Contracts.ETHAJoin
	addrs["eth_b_join"] = cfg.Contracts.ETHBJoin
	addrs["dai_a_join"] = cfg.Contracts.DAIAJoin
	addrs["dai_b_join"] = cfg.Contracts.DAIBJoin
	addrs["dai_median"] = cfg.Contracts.DAIMedian
	addrs["eth_median"] = cfg.Contracts.ETHMedian

	vaultLoader := loaders.NewVaultLoader(
		eth,
		postgresStore,
		addrs["cdp_manager"],
		addrs["get_cdps"],
		addrs["vat"],
	)

	ilksLoader := loaders.NewIlksLoader(
		ceth,
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
		},
		map[common.Address]common.Address{
			cfg.Contracts.DAI:  addrs["dai_median"],
			cfg.Contracts.WETH: addrs["eth_median"],
		},
	)

	ilks, err := ilksLoader.LoadIlks(context.Background())
	if err != nil {
		log.Fatal("error loading ilks.", err)
	}

	var clipperTakeServices []*take.Service
	var clipperRedoServices []*redo.Service

	for _, ilk := range ilks {
		err = postgresStore.CreateOrUpdateIlk(context.Background(), ilk)
		if err != nil {
			log.Fatal("error saving ilks.", err)
		}

		clipperTakeServices = append(
			clipperTakeServices,
			take.NewService(
				take.WithEthClient(eth),
				take.WithClipperAddr(ilk.Clipper),
				take.WithQuoter(quoter),
				take.WithPath(cfg.FindIlkUniswapPath(ilk.Name)),
				take.WithGemJoin(ilk.Join),
				take.WithAssetDecimals(cfg.FindIlkDecimals(ilk.Name)),
				take.WithSender(sender),
				take.WithLogger(logger),
				take.WithIlkName(ilk.Name),
			),
		)

		clipperRedoServices = append(
			clipperRedoServices,
			redo.NewService(
				eth,
				ilk.Clipper,
				sender,
				logger,
			),
		)
	}

	vatLoader := loaders.NewVatLoader(
		eth,
		cfg.Contracts.Vat,
	)

	go func() {
		for {
			for _, rs := range clipperRedoServices {
				err = rs.Start(context.Background())
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	go func() {
		for {
			dogBarkService := bark.NewService(
				context.Background(),
				eth,
				postgresStore,
				addrs["dog"],
				addrs["spot"],
				vaultLoader,
				vatLoader,
				ilksLoader,
				sender,
				logger,
			)

			dogBarkService.Start(context.Background())
		}
	}()

	go func() {
		for {
			for _, ts := range clipperTakeServices {
				err = ts.Start(
					context.Background(),
					big.NewInt(cfg.Processor.MinProfitPercentage),
					big.NewInt(cfg.Processor.MinLotZarValue),
					big.NewInt(cfg.Processor.MaxLotZarValue),
					cfg.Wallet.Address,
				)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	os.Exit(1)
}
