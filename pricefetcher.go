package pricefetcher

import (
	"net/http"
)

const (
	coingecko     = "https://api.coingecko.com/api/v3/simple/price?ids=monero&vs_currencies=%s"
	cryptocompare = "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=XMR&tsyms=%s"
	binance       = "https://api.binance.com/api/v3/ticker/price?symbol=XMRUSDT"
)

var options = []string{"coingecko", "cryptcompare", "binance"}

type Client struct {
	httpcl *http.Client
}

func New(httpClient *http.Client) *Client {
	var cl Client
	if httpClient == nil {
		cl.httpcl = http.DefaultClient
	} else {
		cl.httpcl = httpClient
	}

	return &cl
}
