// Package binance provides a lightweight wrapper for the Binance REST API.
// It is a work in progress.
// This package focuses on endpoints and functions allowed by the binance.us API, and it has not been tested on binance.com.
package binance

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	ApiKey         string
	SecretKey      string
	BaseUrl        BaseUrl
	httpClient     *http.Client
	Symbols        map[string]Symbol
	RequestTimeout time.Duration
}

func NewClient(apiKey, secretKey string, baseUrl BaseUrl) (Client, error) {
	c := Client{
		apiKey,
		secretKey,
		baseUrl,
		http.DefaultClient,
		make(map[string]Symbol),
		2 * time.Second,
	}
	exinfo, err := c.ExchangeInfo()
	if err != nil {
		return c, fmt.Errorf("could not properly initialize client; check internet connection %w", err)
	}
	for _, sym := range exinfo.Symbols {
		c.Symbols[sym.Symbol] = ParseFilters(sym)
	}
	return c, nil
}
