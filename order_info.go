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

type OrderInfoResp struct {
	Symbol                  string  `json:"symbol"`
	OrderId                 int     `json:"orderId"`
	OrderListId             int     `json:"orderListId"` // Unless part of an OCO, the value will always be -1.
	ClientOrderId           string  `json:"clientOrderId"`
	Price                   float64 `json:"price,string"`
	OrigQty                 float64 `json:"origQty,string"`
	ExecutedQty             float64 `json:"executedQty,string"`
	CummulativeQuoteQty     int     `json:"cummulativeQuoteQty,string"`
	Status                  string  `json:"status"`
	TimeInForce             string  `json:"timeInForce"`
	OrderType               string  `json:"type"`
	Side                    string  `json:"side"`
	StopPrice               float64 `json:"stopPrice,string"`
	IcebergQty              float64 `json:"icebergQty,string"`
	Time                    int     `json:"time"`
	UpdateTime              int     `json:"updateTime"`
	IsWorking               bool    `json:"isWorking"`
	OrigQuoteOrderQty       float64 `json:"origQuoteOrderQty,string"`
	WorkingTime             int     `json:"workingTime"`
	SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
}

func (c Client) OrderInfo(symbol string, orderId int) (OrderInfoResp, error) {
	var res OrderInfoResp
	query := "orderId=" + strconv.Itoa(orderId) +
		"&symbol=" + symbol +
		"&timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ORDER + "?" + query + "&signature=" + signature
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
