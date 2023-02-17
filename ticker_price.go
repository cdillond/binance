package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TickerResp struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

func (c Client) TickerPrice(symbol string) (TickerResp, error) {

	var res TickerResp
	if symbol == "" {
		return res, fmt.Errorf("symbol parameter is required; use TickerPrices and an empty slice if you wish to obtain price data for all tickers")
	}
	query := "symbol=" + symbol
	url := string(c.BaseUrl) + TICKER_PRICE + "?" + query
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
		e, err := ParseRespErr(b)
		if err != nil {
			return res, err
		}
		return res, fmt.Errorf("%v %v", e.Code, e.Msg)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c Client) TickerPrices(symbols []string) ([]TickerResp, error) {

	var res []TickerResp
	var query string

	if len(symbols) != 0 {
		b, err := json.Marshal(symbols)
		if err != nil {
			return res, err
		}
		query = "symbols=" + string(b)
	}
	url := string(c.BaseUrl) + TICKER_PRICE + "?" + query
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
		e, err := ParseRespErr(b)
		if err != nil {
			return res, err
		}
		return res, fmt.Errorf("%v %v", e.Code, e.Msg)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
