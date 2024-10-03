package pricefetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Only USDT rate is available
func (c *Client) FetchFromBinance() (float64, error) {
	resp, err := c.httpcl.Get(binance)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Binance API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	priceStr, ok := result["price"]
	if !ok {
		return 0, errors.New("price not found in Binance response")
	}

	var price float64
	fmt.Sscanf(priceStr, "%f", &price)

	return price, nil
}
