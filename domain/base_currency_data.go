package domain

type BaseCurrencyData struct {
	MarketReferenceCurrencyUnit       string `json:"market_reference_currency_unit"`
	MarketReferenceCurrencyPriceInUsd string `json:"market_reference_currency_price_in_usd"`
	NetworkBaseTokenPriceInUsd        string `json:"network_base_token_price_in_usd"`
	NetworkBaseTokenPriceDecimals     int    `json:"network_base_token_price_decimals"`
}
