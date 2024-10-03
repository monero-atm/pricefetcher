package pricefetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type krakenPair struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XmrEur struct {
			C []string `json:"c"`
		} `json:"XXMRZEUR"`
		XmrUsd struct {
			C []string `json:"c"`
		} `json:"XXMRZUSD"`
	} `json:"result"`
}

func (c *Client) FetchFromKraken(currency string) (float64, error) {
	currency = strings.ToUpper(currency)
	url := fmt.Sprintf(kraken, currency)
	resp, err := c.httpcl.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Kraken API error: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var kp krakenPair
	err = json.Unmarshal(body, &kp)
	if err != nil {
		return 0, err
	}
	if len(kp.Error) > 0 {
		return 0, fmt.Errorf("Kraken API returned errors")
	}
	var rate string
	if currency == "EUR" {
		rate = kp.Result.XmrEur.C[0]
	} else if currency == "USD" {
		rate = kp.Result.XmrUsd.C[0]
	} else {
		return 0, fmt.Errorf("non-existent XMR pair on Kraken")
	}
	return strconv.ParseFloat(rate, 64)
}
