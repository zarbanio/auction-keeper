package configs

import (
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Indexer struct {
		BlockInterval time.Duration `yaml:"BlockInterval"`
		PoolSize      int           `yaml:"PoolSize"`
		BlockPtr      uint64        `yaml:"BlockPtr"`
	} `yaml:"General"`
	Redis struct {
		URL string `yaml:"URL"`
	}
	Network struct {
		ChainId     int64 `yaml:"ChainId"`
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
	ZarJoin                common.Address `yaml:"ZarJoin"`
	Vat                    common.Address `yaml:"Vat"`
	UniswapV3QuoterAddress common.Address `yaml:"UniswapV3QuoterAddress"`
	Collaterals            []struct {
		Name            string                    `yaml:"Name"`
		Erc20addr       common.Address            `yaml:"Erc20addr"`
		Decimals        int64                     `yaml:"Decimals"`
		Clipper         common.Address            `yaml:"Clipper"`
		GemJoinAdapter  common.Address            `yaml:"GemJoinAdapter"`
		UniswapV3Callee common.Address            `yaml:"UniswapV3Callee"`
		UniswapV3Path   []entities.UniswapV3Route `yaml:"UniswapV3Path"`
	}
	Processor struct {
		MinProfitPercentage int64 `yaml:"MinProfitPercentage"`
		MinLotZarValue      int64 `yaml:"MinLotZarValue"`
		MaxLotZarValue      int64 `yaml:"MaxLotZarValue"`
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
