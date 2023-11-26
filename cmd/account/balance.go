package account

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/ierc20"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/eoa"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

func balance(cfg configs.Config) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		log.Fatal(err)
	}
	newSigner, err := signer.NewSigner(cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}
	sender := sender.NewSender(newSigner, postgresStore, eth)

	vat, err := vat.NewVat(cfg.Contracts.Vat, eth)
	if err != nil {
		log.Fatal(err)
	}

	zarInSystem, err := vat.Zar(&bind.CallOpts{Context: context.Background()}, sender.GetAddress())
	if err != nil {
		log.Fatal(err)
	}

	erc20s := make(map[string]*ierc20.Ierc20)

	zar, err := ierc20.NewIerc20(cfg.Contracts.ZAR, eth)
	if err != nil {
		log.Fatal(err)
	}

	erc20s["zar"] = zar

	eoa := eoa.NewEOA(sender, erc20s)

	balances, err := eoa.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Address:", sender.GetAddress().String())

	fmt.Println("System balances:")
	b := new(big.Int).Div(zarInSystem, math.BigPow(10, 27))
	v, err := decimal.NewFromString(b.String())
	if err != nil {
		log.Fatal(err)
	}
	s := v.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))).String()
	fmt.Printf("zar: %s\n", s)
	fmt.Println("Wallet balances:")
	for symbol, balance := range balances {
		fmt.Printf("%s: %s\n", symbol, balance.String())
	}
}
