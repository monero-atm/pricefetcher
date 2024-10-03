package pricefetcher

import "errors"

// Returns price, name of the source ("binance", "coingecko" or "cryptocompare") and error if any
func (c *Client) FetchXMRPrice(currency string) (float64, string, error) {
	price, err := c.FetchFromCoinGecko(currency)
	if err == nil {
		return price, "coingecko", nil
	}

	price, err = c.FetchFromCryptoCompare(currency)
	if err == nil {
		return price, "crytocompare", nil
	}

	price, err = c.FetchFromBinance()
	if err == nil {
		return price, "binance", nil
	}

	return 0, "", errors.New("failed to fetch XMR price from all sources")
}