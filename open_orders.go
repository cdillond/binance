package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type rawOpenOrderResp struct {
	Symbol            string  `json:"symbol"`
	OrderId           int     `json:"orderId"`
	OrderListId       int     `json:"orderListId"`
	ClientOrderId     string  `json:"clientOrderId"`
	Price             float64 `json:"price,string"`
	OrigQty           float64 `json:"origQty,string"`
	ExecutedQty       float64 `json:"executedQty,string"`
	CummulativeQty    float64 `json:"cummulativeQuoteQty,string"`
	Status            string  `json:"status"`
	TimeInForce       string  `json:"timeInForce"`
	OrderType         string  `json:"type"`
	Side              string  `json:"side"`
	StopPrice         float64 `json:"stopPrice,string"`
	IcebergQty        float64 `json:"icebergQty,string"`
	Time              int     `json:"time"`
	UpdateTime        int     `json:"updateTime"`
	IsWorking         bool    `json:"isWorking"`
	OrigQuoteOrderQty float64 `json:"origQuoteOrderQty,string"`
	SelfTradeMode     string  `json:"selfTradePreventionMode"`
}

type OpenOrderResp struct {
	Symbol            string
	OrderId           int
	OrderListId       int
	ClientOrderId     string
	Price             float64
	OrigQty           float64
	ExecutedQty       float64
	CummulativeQty    float64
	Status            string
	TimeInForce       TimeInForce
	OrderType         OrderType
	Side              Side
	StopPrice         float64
	IcebergQty        float64
	Time              time.Time
	UpdateTime        time.Time
	IsWorking         bool
	OrigQuoteOrderQty float64
	SelfTradeMode     SelfTradePreventionMode
}

func (c Client) OpenOrders() ([]OpenOrderResp, error) {
	var res []OpenOrderResp
	timestamp := "timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(timestamp)
	url := string(c.BaseUrl) + OPEN_ORDERS + "?" + timestamp + "&signature=" + signature
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
	var raw []rawOpenOrderResp
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return res, parseRespErr(b)
	}

	err = json.Unmarshal(b, &raw)
	if err != nil {
		return res, err
	}
	res = make([]OpenOrderResp, len(raw))

	for i, rr := range raw {
		res[i] = unRaw(rr)
	}
	return res, nil
}

func unRaw(r rawOpenOrderResp) OpenOrderResp {
	return OpenOrderResp{
		Symbol:            r.Symbol,
		OrderId:           r.OrderId,
		OrderListId:       r.OrderListId,
		ClientOrderId:     r.ClientOrderId,
		Price:             r.Price,
		OrigQty:           r.OrigQty,
		ExecutedQty:       r.ExecutedQty,
		CummulativeQty:    r.CummulativeQty,
		Status:            r.Status,
		TimeInForce:       TimeInForce(r.TimeInForce),
		OrderType:         OrderType(r.OrderType),
		Side:              Side(r.Side),
		StopPrice:         r.StopPrice,
		IcebergQty:        r.IcebergQty,
		Time:              time.Unix(0, int64(r.Time)*1_000_000),
		UpdateTime:        time.Unix(0, int64(r.UpdateTime)*1_000_000),
		IsWorking:         r.IsWorking,
		OrigQuoteOrderQty: r.OrigQuoteOrderQty,
		SelfTradeMode:     SelfTradePreventionMode(r.SelfTradeMode),
	}
}
