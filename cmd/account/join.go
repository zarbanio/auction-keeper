package account

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/ierc20"
	"github.com/zarbanio/auction-keeper/bindings/zarban/zarjoin"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

func joinZar(cfg configs.Config, amount *big.Int) {
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
	zarjoin, err := zarjoin.NewZarjoin(cfg.Contracts.ZarJoin, eth)
	if err != nil {
		log.Fatal(err)
	}

	zar, err := ierc20.NewIerc20(cfg.Contracts.ZAR, eth)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := sender.GetTransactOpts()
	if err != nil {
		log.Fatal(err)
	}
	tx, err := zar.Approve(opts, cfg.Contracts.ZarJoin, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("approve tx sent:", tx.Hash().String())
	err = sender.HandleSentTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("approve tx confirmed:", tx.Hash().String())

	ops, err := sender.GetTransactOpts()
	if err != nil {
		log.Fatal(err)
	}
	tx, err = zarjoin.Join(ops, sender.GetAddress(), amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("join tx sent:", tx.Hash().String())
	err = sender.HandleSentTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("join tx confirmed:", tx.Hash().String())
}
