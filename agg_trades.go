package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// letters are assumed to follow order of RecentTradesResp
type AggregateTrade struct {
	A            int64   `json:"a"`
	P            float64 `json:"p,string"`
	Q            float64 `json:"q,string"`
	F            int64   `json:"f"`
	L            int64   `json:"L"`
	T            int64   `json:"T"`
	IsBuyerMaker bool    `json:"m"`
	IsBestMatch  bool    `json:"M"`
}

func (c Client) AggregateTrades(symbol string, limit, fromId int, startTime, endTime time.Time) ([]AggregateTrade, error) {
	q := "?symbol=" + symbol + "&limit=" + strconv.Itoa(limit)

	if fromId >= 0 {
		q = q + "&fromId=" + strconv.Itoa(fromId)
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+AGG_TRADES+q, nil)
	if err != nil {
		return []AggregateTrade{}, err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []AggregateTrade{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []AggregateTrade{}, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return []AggregateTrade{}, parseRespErr(b)
	}

	var tr []AggregateTrade
	err = json.Unmarshal(b, &tr)
	return tr, err
}
