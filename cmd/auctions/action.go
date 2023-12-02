package auctions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/logger"
	"github.com/zarbanio/auction-keeper/services/redo"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

type Mode string

const (
	Bark Mode = "bark"
	Redo Mode = "redo"
	Take Mode = "take"
)

func action(cfg configs.Config, mode Mode, ilkName string, auctionId *big.Int) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	logger := logger.NewLogger(context.Background(), postgresStore)

	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		log.Fatal(err)
	}

	memCache := cache.NewMemCache()

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

	addrs["cdp_manager"] = cfg.Contracts.CDPManager
	addrs["get_cdps"] = cfg.Contracts.GetCDPs
	addrs["ilk_registry"] = cfg.Contracts.IlkRegistry
	addrs["eth_a_join"] = cfg.Contracts.ETHAJoin
	addrs["eth_b_join"] = cfg.Contracts.ETHBJoin
	addrs["dai_a_join"] = cfg.Contracts.DAIAJoin
	addrs["dai_b_join"] = cfg.Contracts.DAIBJoin
	addrs["dai_median"] = cfg.Contracts.DAIMedian
	addrs["eth_median"] = cfg.Contracts.ETHMedian
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

	for _, ilk := range ilks {
		if ilkName == ilk.Name {
			clipper, err := clipper.NewClipper(ilk.Clipper, eth)
			if err != nil {
				log.Fatal(err)
			}
			list, err := clipper.List(&bind.CallOpts{Context: context.Background()})
			if err != nil {
				log.Fatal(err)
			}
			for _, id := range list {
				if !intToBool(id.Cmp(auctionId)) { // that is, id == auctionId
					if mode == "redo" {
						redoService := redo.NewService(ilk.Name, eth, ilk.Clipper, sender, logger)
						err := redoService.Redo(id)
						if err != nil {
							logger.ConsoleLogger.Err(err).Msg("error while redoing the auction.")
						}
					}
					if mode == "take" {
						// ! TODO: TAKE SERVICE
						fmt.Println("will be implemented")
					}
					break
				}
			}
			break
		}

	}
}

func intToBool(n int) bool {
	return n != 0
}
