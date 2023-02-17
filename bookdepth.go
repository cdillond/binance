package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type tmpBookDepth struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type BookDepth struct {
	LastUpdateId int
	Bids         [][]float64
	Asks         [][]float64
}

func (c Client) OrderBookDepth(symbol string, limit int) (BookDepth, error) {
	var res BookDepth
	q := "?symbol=" + symbol + "&limit=" + strconv.Itoa(limit)

	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+BOOK_DEPTH+q, nil)
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
	var tr tmpBookDepth
	err = json.Unmarshal(b, &tr)
	if err != nil {
		return res, err
	}
	res.LastUpdateId = tr.LastUpdateId

	res.Bids, res.Asks = make([][]float64, len(tr.Bids)), make([][]float64, len(tr.Bids))
	for i := 0; i < len(tr.Bids); i++ {
		a, err := strconv.ParseFloat(tr.Bids[i][0], 64)
		if err != nil {
			return res, err
		}
		b, err := strconv.ParseFloat(tr.Bids[i][1], 64)
		if err != nil {
			return res, err
		}
		res.Bids[i] = []float64{a, b}
	}
	for i := 0; i < len(tr.Asks); i++ {
		a, err := strconv.ParseFloat(tr.Asks[i][0], 64)
		if err != nil {
			return res, err
		}
		b, err := strconv.ParseFloat(tr.Asks[i][1], 64)
		if err != nil {
			return res, err
		}
		res.Asks[i] = []float64{a, b}
	}
	return res, err
}
