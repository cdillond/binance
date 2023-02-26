package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PriceChange struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	BidQty             float64 `json:"bidQty,string"`
	AskPrice           float64 `json:"askPrice,string"`
	AskQty             float64 `json:"askQty,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int     `json:"openTime"`
	CloseTime          int     `json:"closeTime"`
	FirstId            int     `json:"firstId"`
	LastId             int     `json:"lastId"`
	Count              int     `json:"count"`
}

// Returns the 24-hour price change statistics
func (c Client) DayPriceChange(symbol string) (PriceChange, error) {
	var res PriceChange
	if symbol == "" {
		return res, fmt.Errorf("symbol parameter is required; use DayPriceChanges and an empty slice if you wish to obtain price data for all tickers")
	}
	query := "symbol=" + symbol
	url := string(c.BaseUrl) + DAY_CHANGE + "?" + query
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return res, parseRespErr(b)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Returns the 24-hour price change statistics
func (c Client) DayPriceChanges(symbols []string) ([]PriceChange, error) {
	var res []PriceChange
	var query string
	if len(symbols) != 0 {
		s, err := json.Marshal(symbols)
		if err != nil {
			return res, err
		}
		query = "symbols=" + string(s)
	}

	url := string(c.BaseUrl) + DAY_CHANGE + "?" + query
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return res, parseRespErr(b)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
