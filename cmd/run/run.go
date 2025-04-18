package run

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/bark"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/logger"
	"github.com/zarbanio/auction-keeper/services/redo"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/services/take"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
	"github.com/zarbanio/auction-keeper/store"
)

type Mode string

const (
	Bark Mode = "bark"
	Redo Mode = "redo"
	Take Mode = "take"
)

func main(cfg configs.Config, secrets configs.Secrets, modes []Mode, useUniswap bool, allowedIlks []string) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	logger := logger.NewLogger(context.Background(), postgresStore)

	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}

	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	memCache := cache.NewMemCache()

	newSigner, err := signer.NewSigner(secrets.PrivateKey, big.NewInt(cfg.Network.ChainId))
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

	vaultLoader := loaders.NewVaultLoader(
		eth,
		postgresStore,
		cfg.Contracts.CDPManager,
		addrs["vat"],
	)

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
		log.Fatal("error loading ilks.", err)
	}

	var clipperTakeServices []*take.Service
	var clipperRedoServices []*redo.Service

	for _, ilk := range ilks {
		err = postgresStore.CreateOrUpdateIlk(context.Background(), ilk)
		if err != nil {
			log.Fatal("error saving ilks.", err)
		}

		if len(allowedIlks) > 0 && !contains(allowedIlks, ilk.Name) {
			continue
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
				take.WithCallee(cfg.Contracts.UniswapV3Callee),
				take.WithUseUniswap(useUniswap),
			),
		)

		clipperRedoServices = append(
			clipperRedoServices,
			redo.NewService(
				ilk.Name,
				eth,
				ilk.Clipper,
				sender,
				logger,
			),
		)
	}

	dogBarkService := bark.NewService(
		context.Background(),
		eth,
		postgresStore,
		addrs["dog"],
		addrs["spot"],
		vaultLoader,
		ilksLoader,
		sender,
		logger,
	)

	for _, mode := range modes {
		switch mode {
		case Bark:
			go func() {
				for {
					dogBarkService.Start(context.Background())
					time.Sleep(cfg.Times.BarkTicker)
				}
			}()
		case Redo:
			go func() {
				for {
					for _, rs := range clipperRedoServices {
						err = rs.Run(context.Background())
						if err != nil {
							log.Fatal(err)
						}
					}
					time.Sleep(cfg.Times.RedoTicker)
				}
			}()
		case Take:
			go func() {
				for {
					for _, ts := range clipperTakeServices {
						err = ts.Start(
							context.Background(),
							big.NewInt(cfg.Processor.MinProfitPercentage),
							big.NewInt(cfg.Processor.MinLotZarValue),
							big.NewInt(cfg.Processor.MaxLotZarValue),
							secrets.WalletAddress,
						)
						if err != nil {
							log.Fatal(err)
						}
					}
					time.Sleep(cfg.Times.TakeTicker)
				}
			}()
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func Register(root *cobra.Command) {
	root.PersistentFlags().String("modes", "bark,redo,take", "run mode")
	root.PersistentFlags().String("ilks", "daia,daib,etha,ethb,wsteth-a", "ilks to run on")
	root.PersistentFlags().Bool("uniswap", false, "set it to true for swapping the released collateral to zar for liquidating a vault")
	root.AddCommand(
		&cobra.Command{
			Use:   "run",
			Short: "run keeper",
			Run: func(cmd *cobra.Command, args []string) {
				configFile, _ := cmd.Flags().GetString("config")
				mode, _ := cmd.Flags().GetString("modes")
				useUniswap, _ := cmd.Flags().GetBool("uniswap")
				tokens := strings.Split(mode, ",")

				if len(tokens) == 0 {
					log.Fatal("no mode specified")
				}

				modes := []Mode{}
				for _, m := range tokens {
					if m != string(Bark) && m != string(Redo) && m != string(Take) {
						log.Fatal("invalid mode specified")
					}
					modes = append(modes, Mode(m))
				}

				ilks, _ := cmd.Flags().GetString("ilks")
				allowedIlks := strings.Split(ilks, ",")

				cfg := configs.ReadConfig(configFile)
				secrets := configs.ReadSecrets()
				main(cfg, secrets, modes, useUniswap, allowedIlks)
			},
		},
	)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
