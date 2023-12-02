package auctions

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/domain/math"
)

func Register(root *cobra.Command) {
	auctions := &cobra.Command{
		Use:   "auctions",
		Short: "auctions, a command to manage auctions. see clippers, etc",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("no command specified. see help for more info. ex: auctions clippers")
		},
	}
	auctions.AddCommand(&cobra.Command{
		Use: "clippers",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)

			listClippers(cfg)
		},
	})
	auctions.AddCommand(&cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			listAuctions(cfg)
		},
	})
	redoCmd := &cobra.Command{
		Use: "redo [ilkName] [auctionId]",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			// Check if ID is provided as an argument
			if len(args) > 1 {
				ilkName := strings.ToUpper(args[0])
				auctionId := math.BigIntFromString(args[1])
				fmt.Printf("Performing redo on ilk: %s with auction ID: %s\n", ilkName, auctionId)
				action(cfg, "redo", ilkName, auctionId)
			} else {
				fmt.Println("Not enough args! Usage: redo [ilkName] [auctionId]")
			}

		},
	}

	auctions.AddCommand(redoCmd)
	root.AddCommand(auctions)
}
