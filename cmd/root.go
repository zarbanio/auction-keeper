package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/cmd/account"
	"github.com/zarbanio/auction-keeper/cmd/auctions"
	"github.com/zarbanio/auction-keeper/cmd/run"
)

func Execute() {
	root := &cobra.Command{
		Use:     "keeper",
		Short:   "api for zarban",
		Version: "0.1",
	}

	root.PersistentFlags().String("config", "config.goerli.yaml", "read config file")
	run.Register(root)
	account.Register(root)
	auctions.Register(root)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
