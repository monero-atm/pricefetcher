# Pricefetcher

A simple library for fetching Monero rate for USD or EUR.

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/monero-atm/pricefetcher"
)

func main() {
	// Pass nil as an argument to use the default http.Client
	client := pricefetcher.New(nil)

	price, err := client.FetchFromCoinGecko("EUR")
	if err != nil {
		log.Println("Failed to fetch price from coingecko:", err)
	} else {
		fmt.Printf("CoinGecko EUR price = %f\n", price)
	}

	price, err = client.FetchFromCryptoCompare("EUR")
	if err != nil {
		log.Println("Failed to fetch price from crytocompare:", err)
	} else {
		fmt.Printf("CryptoCompare EUR price = %f\n", price)
	}

	// Binance only supports USDT. To convert that to EUR I suggest getting EUR/USD rate from ECB.
	price, err = client.FetchFromBinance()
	if err != nil {
		log.Println("Failed to fetch price from binance:", err)
	} else {
		fmt.Printf("Binance USDT price = %f\n", price)
	}

	price, source, err := client.FetchXMRPrice("USD")
	if err != nil {
		log.Println("Failed to get rate with fallbacks:", err)
	} else {
		fmt.Printf("Got USD rate=%f from source=%s\n", price, source)
	}
}
```
