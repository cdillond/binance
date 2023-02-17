package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Default limit is 500; value must be between 1 and 1,000. If fromId is negative, the most recent trades are used.
func (c *Client) HistoricalTrades(symbol string, limit, fromId int) ([]RecentTradesResp, error) {
	q := "?symbol=" + symbol + "&limit=" + strconv.Itoa(limit)

	if fromId >= 0 {
		q = q + "&fromId=" + strconv.Itoa(fromId)
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+HIST_TRADES+q, nil)
	if err != nil {
		return []RecentTradesResp{}, err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
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
