package account

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/spf13/cobra"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

func Register(root *cobra.Command) {
	account := &cobra.Command{
		Use:   "account",
		Short: "account, a command to manage account. see balance, allowance, approve, etc",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("no command specified. see help for more info. ex: account balance")
		},
	}
	account.AddCommand(&cobra.Command{
		Use: "balance",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			secrets := configs.ReadSecrets()
			balance(cfg, secrets)
		},
	})
	account.AddCommand(&cobra.Command{
		Use: "hope",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			secrets := configs.ReadSecrets()

			if len(args) == 0 {
				fmt.Println("you must specify at least one address")
				fmt.Println("ex: account hope 0x1234 0x5678")
				return
			}
			addrs := make([]common.Address, 0, len(args))
			for _, arg := range args {
				addrs = append(addrs, common.HexToAddress(arg))
			}
			hope(cfg, secrets, addrs...)
		},
	})
	account.AddCommand(&cobra.Command{
		Use: "join",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			cfg := configs.ReadConfig(configFile)
			secrets := configs.ReadSecrets()

			if len(args) == 0 {
				fmt.Println("you must specify amount to join")
				fmt.Println("ex: account join 100.15")
				return
			}

			amount, err := decimal.NewFromString(args[0])
			if err != nil {
				fmt.Println("invalid amount specified", err)
				return
			}
			amount = amount.Mul(decimal.NewFromInt(1000))
			b, err := amount.BigInt()
			if err != nil {
				fmt.Println("invalid amount specified", err)
				return
			}
			b = new(big.Int).Mul(math.BigPow(10, 15), b)
			joinZar(cfg, secrets, b)
		},
	})
	root.AddCommand(account)
}
