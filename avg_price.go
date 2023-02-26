package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AvgPriceResp struct {
	Mins  int     `json:"mins"`
	Price float64 `json:"price,string"`
}

func (c Client) AvgPrice(symbol string) (AvgPriceResp, error) {

	var res AvgPriceResp
	if symbol == "" {
		return res, fmt.Errorf("symbol parameter is required; use TickerPrices and an empty slice if you wish to obtain price data for all tickers")
	}
	query := "symbol=" + symbol
	url := string(c.BaseUrl) + AVG_PRICE + "?" + query
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
