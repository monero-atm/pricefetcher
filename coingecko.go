package pricefetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// EUR and USD rates are available
func (c *Client) FetchFromCoinGecko(currency string) (float64, error) {
	url := fmt.Sprintf(coingecko, currency)
	resp, err := c.httpcl.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("CoinGecko API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]map[string]float64
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	price, ok := result["monero"][strings.ToLower(currency)]
	if !ok {
		return 0, errors.New("currency not found in CoinGecko response")
	}

	return price, nil
}
