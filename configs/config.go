package configs

import (
	"log"
	"math/big"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/entities"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Indexer struct {
		BlockInterval time.Duration `yaml:"BlockInterval"`
		StartBlock    uint64        `yaml:"StartBlock"`
		BlockRange    int64         `yaml:"BlockRange"`
	} `yaml:"General"`
	Redis struct {
		URL string `yaml:"URL"`
	}
	Postgres struct {
		Host           string `yaml:"Host"`
		User           string `yaml:"User"`
		Password       string `yaml:"Password"`
		DB             string `yaml:"DB"`
		MigrationsPath string `yaml:"MigrationsPath"`
	} `yaml:"Postgres"`
	Network struct {
		Name        string `yaml:"Name"`
		ChainId     int64  `yaml:"ChainId"`
		NativeAsset struct {
			Name        string         `yaml:"Name"`
			Symbol      string         `yaml:"Symbol"`
			Decimals    int64          `yaml:"Decimals"`
			MockAddress common.Address `yaml:"MockAddress"`
		}
		Node struct {
			Api string `yaml:"Api"`
		}
	}
	Wallet struct {
		Private string         `yaml:"Private"`
		Address common.Address `yaml:"Address"`
	}
	Ilks []struct {
		Name          string                    `yaml:"Name"`
		Decimals      int64                     `yaml:"Decimals"`
		UniswapV3Path []entities.UniswapV3Route `yaml:"UniswapV3Path"`
	} `yaml:"Ilks"`
	Processor struct {
		MinProfitPercentage int64 `yaml:"MinProfitPercentage"`
		MinLotZarValue      int64 `yaml:"MinLotZarValue"`
		MaxLotZarValue      int64 `yaml:"MaxLotZarValue"`
	}
	Times struct {
		VaultTicker      int64 `yaml:"VaultTicker"`
		FlopperTicker    int64 `yaml:"FlopperTicker"`
		LiquidatorTicker int64 `yaml:"LiquidatorTicker"`
	}
	Contracts struct {
		Deployment      common.Address `yaml:"Deployment"`
		IlkRegistry     common.Address `yaml:"IlkRegistry"`
		AddressProvider common.Address `yaml:"AddressProvider"`
		CDPManager      common.Address `yaml:"CDPManager"`
		GetCDPs         common.Address `yaml:"GetCDPs"`
		ETHAJoin        common.Address `yaml:"ETHAJoin"`
		ETHBJoin        common.Address `yaml:"ETHBJoin"`
		DAIAJoin        common.Address `yaml:"DAIAJoin"`
		DAIBJoin        common.Address `yaml:"DAIBJoin"`
		DAIMedian       common.Address `yaml:"DAIMedian"`
		ETHMedian       common.Address `yaml:"ETHMedian"`
		DAI             common.Address `yaml:"DAI"`
		ZAR             common.Address `yaml:"ZAR"`
		WETH            common.Address `yaml:"WETH"`
		UniswapV3Quoter common.Address `yaml:"UniswapV3Quoter"`
		UniswapV3Callee common.Address `yaml:"UniswapV3Callee"`
		ZarJoin         common.Address `yaml:"ZarJoin"`
		Vat             common.Address `yaml:"Vat"`
		Vow             common.Address `yaml:"Vow"`
		Dog             common.Address `yaml:"Dog"`
	}
}

func ReadConfig(configFile string) Config {
	c := &Config{}
	err := c.Unmarshal(c, configFile)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return *c
}

func (c *Config) Unmarshal(rawVal interface{}, fileName string) error {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	var input interface{} = viper.AllSettings()
	config := defaultDecoderConfig(rawVal)
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func defaultDecoderConfig(output interface{}) *mapstructure.DecoderConfig {
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			EthAddressDecodeHookFunc(),
		),
	}
	return c
}

func EthAddressDecodeHookFunc() mapstructure.DecodeHookFuncType {
	// Wrapped in a function call to add optional input parameters (eg. separator)
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		// Check if the data type matches the expected one
		if f.Kind() != reflect.String {
			return data, nil
		}

		// Check if the target type matches the expected one
		if t != reflect.TypeOf(common.Address{}) {
			return data, nil
		}

		// Format/decode/parse the data and return the new value
		return common.HexToAddress(data.(string)), nil
	}
}

func (c Config) FindIlkUniswapPath(name string) []entities.UniswapV3Route {
	for _, ilk := range c.Ilks {
		if ilk.Name == name {
			return ilk.UniswapV3Path
		}
	}
	return nil
}

func (c Config) FindIlkDecimals(name string) *big.Int {
	for _, ilk := range c.Ilks {
		if ilk.Name == name {
			return big.NewInt(ilk.Decimals)
		}
	}
	return big.NewInt(0)
}
