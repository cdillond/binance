package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c Client) CancelAll(symbol string) ([]TradeAck, error) {
	var res []TradeAck
	query := "symbol=" + symbol +
		"&timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + OPEN_ORDERS + "?" + query + "&signature=" + signature
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return res, err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
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
	//res = make([]TradeResp, len(rawres))
	//for i, tf := range rawres {
	//	res[i] = TradeResp{rawResp: &tf}
	//}
	return res, nil
}
