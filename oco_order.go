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

type OCOOrder struct {
	Symbol                  string // REQUIRED
	ListClientOrderId       string
	Side                    Side    // REQUIRED
	Quantity                float64 // REQUIRED
	LimitClientOrderId      string
	LimitStrategyId         int
	LimitStrategyType       int
	Price                   float64 // REQUIRED
	LimitIcebergQty         float64
	TrailingDelta           int
	StopClientOrderId       string
	StopPrice               float64 // REQUIRED
	StopStrategyId          int
	StopStrategyType        int
	StopLimitPrice          float64
	StopIcebergQty          float64
	StopLimitTimeInForce    TimeInForce // REQUIRED WITH StopLimitPrice
	NewOrderRespType        OrderRespType
	SelfTradePreventionMode SelfTradePreventionMode
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

type OCOTradeResp struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int    `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []TradeFull `json:"orderReports"`
}

func NewOCOOrder(symbol string, side Side, quantity, price, stopPrice, StopLimitPrice float64) OCOOrder {
	return OCOOrder{
		Symbol:               symbol,
		Side:                 side,
		Quantity:             quantity,
		Price:                price,
		StopPrice:            stopPrice,
		StopLimitPrice:       StopLimitPrice,
		StopLimitTimeInForce: GTC,
		Timestamp:            time.Now(),
	}
}

func (o OCOOrder) submitOCO(c *Client) (string, error) {
	var query string
	pair, ok := c.Symbols[o.Symbol]
	if !ok {
		return query, fmt.Errorf("invalid symbol %v or improperly initialized client", o.Symbol)
	}
	var listCoid, limitCoid, limitIcebergQty, trailingDelta,
		stopCoid, stopLimitPrice, stopLimitTIF, nort, stpm string
	if o.ListClientOrderId != "" {
		listCoid = "&listClientOrderId=" + o.ListClientOrderId
	}
	if o.LimitClientOrderId != "" {
		limitCoid = "&limitClientOrderId=" + o.LimitClientOrderId
	}
	if o.LimitIcebergQty != 0 {
		limitIcebergQty = "&limitIcebergQty=" + strconv.FormatFloat(o.LimitIcebergQty, 'f', pair.StepSize, 64)
	}
	if o.TrailingDelta != 0 {
		trailingDelta = "&trailingDelta=" + strconv.Itoa(o.TrailingDelta)
	}
	if o.StopClientOrderId != "" {
		stopCoid = "&stopClientOrderId=" + o.StopClientOrderId
	}
	if o.StopLimitPrice != 0 {
		stopLimitPrice = "&stopLimitPrice=" + strconv.FormatFloat(o.StopLimitPrice, 'f', pair.TickSize, 64)
	}
	if o.StopLimitTimeInForce != "" {
		stopLimitTIF = "&stopLimitTimeInForce=" + string(o.StopLimitTimeInForce)
	}
	if o.NewOrderRespType != "" {
		nort = "&newOrderRespType=" + string(o.NewOrderRespType)
	}
	if o.SelfTradePreventionMode != "" {
		stpm = "&selfTradePreventionMode=" + string(o.SelfTradePreventionMode)
	}

	query = "symbol=" + o.Symbol +
		"&side=" + string(o.Side) +
		"&quantity=" + strconv.FormatFloat(o.Quantity, 'f', pair.StepSize, 64) +
		"&price=" + strconv.FormatFloat(o.Price, 'f', pair.TickSize, 64) +
		"&stopPrice=" + strconv.FormatFloat(o.StopPrice, 'f', pair.TickSize, 64) +
		listCoid + limitCoid + limitIcebergQty + trailingDelta + stopCoid +
		stopLimitPrice + stopLimitTIF + nort + stpm +
		"&timestamp=" + strconv.Itoa(int(o.Timestamp.UnixMilli()))
	return query, nil
}

func (c Client) TradeOCO(o OCOOrder) (OCOTradeResp, error) {
	var res OCOTradeResp
	query, err := o.submitOCO(&c)
	signature := c.Sign(query)
	fmt.Println(query)
	url := string(c.BaseUrl) + OCO_ORDER +
		"?" + query +
		"&signature=" + signature
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
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
	return res, err

}
