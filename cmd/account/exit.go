package account

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/gemjoin"
	"github.com/zarbanio/auction-keeper/bindings/zarban/zarjoin"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

func exitZar(cfg configs.Config, secrets configs.Secrets, amount *big.Int) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}
	defer eth.Close()
	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}
	newSigner, err := signer.NewSigner(secrets.PrivateKey, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}
	sender := sender.NewSender(newSigner, postgresStore, eth)

	zarjoin, err := zarjoin.NewZarjoin(cfg.Contracts.ZarJoin, eth)
	if err != nil {
		log.Fatal(err)
	}

	ops, err := sender.GetTransactOpts()
	if err != nil {
		log.Fatal(err)
	}
	tx, err := zarjoin.Exit(ops, sender.GetAddress(), amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exit zar tx sent:", tx.Hash().String())
	err = sender.HandleSentTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exit zar tx confirmed.")
}

func exitGem(cfg configs.Config, secrets configs.Secrets, ilkName string, amount *big.Int) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	eth, err := ethclient.Dial(secrets.RpcArbitrum)
	if err != nil {
		log.Fatal(err)
	}
	defer eth.Close()
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

	ilk, err := ilksLoader.LoadIlkByName(context.Background(), ilkName)
	if err != nil || ilk.Join == common.HexToAddress("0x") {
		log.Fatal("error fetching ilk.", err)
	}

	gemJoin, err := gemjoin.NewGemjoin(ilk.Join, eth)
	if err != nil {
		log.Fatal(err)
	}

	ops, err := sender.GetTransactOpts()
	if err != nil {
		log.Fatal(err)
	}
	tx, err := gemJoin.Exit(ops, sender.GetAddress(), amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exit gem tx sent:", tx.Hash().String())
	err = sender.HandleSentTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exit gem tx confirmed.")
}
