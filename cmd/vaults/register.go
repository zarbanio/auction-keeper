package vaults

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/domain/math"
)

func Register(root *cobra.Command) {
	vaults := &cobra.Command{
		Use:   "vaults",
		Short: "vaults, a command to see vaults. see clippers, etc",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("no command specified. see help for more info. ex: auctions clippers")
		},
	}

	lsCmd := &cobra.Command{
		Use: "ls [id]",
		Run: func(cmd *cobra.Command, args []string) {
			all, _ := cmd.Flags().GetBool("all")
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			// Check if ID is provided as an argument
			if len(args) > 0 {
				id := args[0]
				fmt.Printf("Listing vault with ID: %s\n", id)
				listVault(cfg, math.BigIntFromString(id))
			} else if all {
				fmt.Println("Listing all vaults")
				listVaults(cfg)
			} else {
				fmt.Println("Listing default set of vaults")
			}

		},
	}

	// Add the --all flag to ls command
	lsCmd.Flags().Bool("all", false, "List all vaults")

	vaults.AddCommand(lsCmd)
	root.AddCommand(vaults)
}
