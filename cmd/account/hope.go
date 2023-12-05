package account

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/bindings/zarban/zarjoin"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

func hope(cfg configs.Config, addrs ...common.Address) {
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		log.Fatal(err)
	}
	defer eth.Close()
	newSigner, err := signer.NewSigner(cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
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
		tx, err := vat.Hope(ops, addr)
		if err != nil {
			log.Fatal(err)
		}
		err = sender.HandleSentTx(tx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func exitZars(zarjoin *zarjoin.Zarjoin, sender sender.Sender, amount *big.Int) {
	ops, err := sender.GetTransactOpts()
	if err != nil {
		log.Fatal(err)
	}
	tx, err := zarjoin.Exit(ops, sender.GetAddress(), amount)
	if err != nil {
		log.Fatal(err)
	}
	err = sender.HandleSentTx(tx)
	if err != nil {
		log.Fatal(err)
	}
}
