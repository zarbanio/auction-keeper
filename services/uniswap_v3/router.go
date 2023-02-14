package uniswap_v3

import (
	"encoding/hex"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"strings"
)

func GetRouter(routes []entities.UniswapV3Route) ([]byte, error) {

	routeStr := ""
	preDestinationToken := ""
	for _, r := range routes {
		fee := fmt.Sprintf("%06x", r.Fee)

		if preDestinationToken == r.TokenA.String() {
			routeStr += fee
			routeStr += strings.Replace(r.TokenB.String(), "0x", "", 1)
		} else {
			routeStr += strings.Replace(r.TokenA.String(), "0x", "", 1)
			routeStr += fee
			routeStr += strings.Replace(r.TokenB.String(), "0x", "", 1)
		}

		preDestinationToken = r.TokenB.String()
	}

	route, err := hex.DecodeString(strings.ToLower(routeStr))
	if err != nil {
		return nil, err
	}

	return route, nil
}
