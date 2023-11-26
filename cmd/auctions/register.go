package auctions

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/configs"
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
	root.AddCommand(auctions)
}
