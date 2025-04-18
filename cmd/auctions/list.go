package auctions

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

func listAuctions(cfg configs.Config, secrets configs.Secrets) {
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

	fmt.Println("Clippers Addresses:")
	clippers := make(map[string]*clipper.Clipper)
	for _, ilk := range ilks {
		fmt.Printf("	%s: %s \n", ilk.Name, ilk.Clipper.String())
		clipper, err := clipper.NewClipper(ilk.Clipper, eth)
		if err != nil {
			log.Fatal(err)
		}
		clippers[ilk.Name] = clipper
	}
	fmt.Println("--------------------------------------------------")
	for ilk, clipper := range clippers {
		list, err := clipper.List(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			log.Fatal(err)
		}
		for _, id := range list {
			sale, err := clipper.Sales(&bind.CallOpts{Context: context.Background()}, id)
			if err != nil {
				log.Fatal(err)
			}
			status, err := clipper.GetStatus(&bind.CallOpts{Context: context.Background()}, id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Sale %s: %s \n", ilk, id.String())
			fmt.Println("	Needs Redo:                       ", status.NeedsRedo)
			fmt.Println("	Liquidated CDP:                   ", sale.Usr.String())
			fmt.Println("	Initial Price Of Auction:         ", Normalize(sale.Top, 27))
			fmt.Println("	Current Price Of Auction:         ", Normalize(status.Price, 27))
			fmt.Println("	Target Zar To Raise:              ", Normalize(sale.Tab, 45))
			fmt.Println("	Collateral Avaiable For Purchase: ", Normalize(sale.Lot, 18))
			fmt.Println("	Auction Start Time:               ", time.Unix(sale.Tic.Int64(), 0).Local().String())
			fmt.Println("--------------------------------------------------")
		}
	}
}

var c map[int64]*big.Int
var once sync.Once

func Pow10(n int64) *big.Int {
	once.Do(func() {
		c = make(map[int64]*big.Int)
	})
	if v, ok := c[n]; ok {
		return v
	}
	c[n] = math.BigPow(10, n)
	return c[n]
}

func Normalize(n *big.Int, decimals int64) decimal.Decimal {
	a := decimal.NewFromBigInt(n)
	b := decimal.NewFromBigInt(Pow10(decimals))
	return a.Div(b)
}
