package router

import (
	"encoding/hex"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"strings"
)

func GetUniswapV3Router(routes []entities.UniswapRoute) ([]byte, error) {

	routeStr := ""
	preDestinationToken := ""
	for _, r := range routes {
		fee := fmt.Sprintf("%06x", r.Fee)

		if preDestinationToken == r.TokenA {
			routeStr += fee
			routeStr += strings.Replace(r.TokenB, "0x", "", 1)
		} else {
			routeStr += strings.Replace(r.TokenA, "0x", "", 1)
			routeStr += fee
			routeStr += strings.Replace(r.TokenB, "0x", "", 1)
		}

		preDestinationToken = r.TokenB
	}

	route, err := hex.DecodeString(strings.ToLower(routeStr))
	if err != nil {
		return nil, err
	}

	return route, nil
}
