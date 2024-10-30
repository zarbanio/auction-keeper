package configs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	envconf "github.com/kelseyhightower/envconfig"
)

type Secrets struct {
	RpcArbitrum   string         `envconfig:"RPC_ARBITRUM" required:"true"`
	PrivateKey    string         `envconfig:"PRIVATE_KEY" required:"true"`
	WalletAddress common.Address `envconfig:"WALLET_ADDRESS" required:"true"`
}

func ReadSecrets() Secrets {
	_ = godotenv.Load()
	var cfg Secrets
	err := envconf.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
