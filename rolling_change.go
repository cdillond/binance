package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c Client) RollingChange(symbol string, windowSize uint, windowSizeUnit IntervalUnit) (PriceChange, error) {
	var res PriceChange
	var ok bool
	var window string
	switch windowSizeUnit {
	case "m":
		if windowSize >= 1 && windowSize <= 59 {
			ok = true
			window = strconv.Itoa(int(windowSize)) + string(windowSizeUnit)
		}
	case "h":
		if windowSize >= 1 && windowSize <= 23 {
			ok = true
			window = strconv.Itoa(int(windowSize)) + string(windowSizeUnit)
		}
	case "d":
		if windowSize >= 1 && windowSize <= 7 {
			ok = true
			window = strconv.Itoa(int(windowSize)) + string(windowSizeUnit)
		}
	}
	if !ok {
		return res, fmt.Errorf("invalid windowSize and windowSizeUnit combination")
	}

	if symbol == "" {
		return res, fmt.Errorf("symbol parameter is required; use RollingChanges and an empty slice if you wish to obtain price data for all tickers")
	}
	query := "symbol=" + symbol +
		"&windowSize=" + window

	url := string(c.BaseUrl) + TICKER + "?" + query
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
