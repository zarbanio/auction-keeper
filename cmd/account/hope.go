package account

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

func hope(cfg configs.Config, secrets configs.Secrets, addrs ...common.Address) {
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
	vat, err := vat.NewVat(cfg.Contracts.Vat, eth)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		ops, err := sender.GetTransactOpts()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hoping for:", addr.String())
		tx, err := vat.Hope(ops, addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hope tx sent:", tx.Hash().String())
		err = sender.HandleSentTx(tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hope tx confirmed.")
	}
}
