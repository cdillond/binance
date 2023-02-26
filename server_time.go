package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type serverTimeResp struct {
	ServerTime int `json:"serverTime"`
}

func (c Client) ServerTime() (time.Time, error) {
	var res time.Time
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+TIME, nil)
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

	var st serverTimeResp
	err = json.Unmarshal(b, &st)
	if err != nil {
		return res, err
	}
	return time.Unix(0, 1000000*int64(st.ServerTime)), nil
}
