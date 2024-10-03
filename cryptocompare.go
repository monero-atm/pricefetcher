package pricefetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type cryptoCompareResponse struct {
	RAW struct {
		XMR map[string]struct {
			Price float64 `json:"PRICE"`
		} `json:"XMR"`
	} `json:"RAW"`
}

// EUR and USD rates are available
func (c *Client) FetchFromCryptoCompare(currency string) (float64, error) {
	currency = strings.ToUpper(currency)
	url := fmt.Sprintf(cryptocompare, currency)
	resp, err := c.httpcl.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("CryptoCompare API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Unmarshal the response into a struct
	var result cryptoCompareResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	// Extract the price
	price, ok := result.RAW.XMR[currency]
	if !ok {
		return 0, errors.New("currency not found in CryptoCompare response")
	}

	return price.Price, nil
}
