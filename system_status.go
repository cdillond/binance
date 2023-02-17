package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type systemStatusResp struct {
	Status int `json:"status"`
}

func (c Client) SystemStatus() (int, error) {
	var status int
	t := strconv.Itoa(int(time.Now().UnixMilli()))
	q := "timestamp=" + t + "&signature=" + c.Sign("timestamp="+t)

	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+SYSTEM_STATUS+"?"+q, nil)
	if err != nil {
		return status, err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return status, err
	}
	defer resp.Body.Close()
	var s systemStatusResp
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return status, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		e, err := ParseRespErr(b)
		if err != nil {
			return status, err
		}
		return status, fmt.Errorf("%v %v", e.Code, e.Msg)
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		return status, err
	}
	return s.Status, nil
}
