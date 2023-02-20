package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type RecentTradesResp struct {
	Id           int     `json:"id"`
	Price        float64 `json:"price,string"`
	Qty          float64 `json:"qty,string"`
	QuoteQty     float64 `json:"quoteQty,string"`
	Time         int64   `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}

// Default limit is 500; value must be between 1 and 1,000
func (c Client) RecentTrades(symbol string, limit int) ([]RecentTradesResp, error) {
	q := "?symbol=" + symbol + "&limit=" + strconv.Itoa(limit)
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+RECENT_TRADES+q, nil)
	if err != nil {
		return []RecentTradesResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []RecentTradesResp{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []RecentTradesResp{}, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		e, err := ParseRespErr(b)
		if err != nil {
			return []RecentTradesResp{}, err
		}
		return []RecentTradesResp{}, fmt.Errorf("%v %v", e.Code, e.Msg)
	}
	var tr []RecentTradesResp
	err = json.Unmarshal(b, &tr)
	return tr, err
}
