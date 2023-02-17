package binance

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ErrResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// The Binance API responds to order creation requests with an error or at most one of an ACK, a RESULT, or a FULL response type, depending on the order type or
// NewOrderRespType parameter. The order types ascend in size and provide increasingly detailed information. A FULL response contains all the information in a RESULT
// response, and a RESULT response contains all the information in an ACK response. A TradeAck and TradeResult can be derived from any valid TradeResp with a RespType of "FULL", and soforth.
// Per the Binance API documentation, "MARKET and LIMIT order types default to FULL; all other orders default to ACK".
type TradeResp struct {
	rawResp *TradeFull
}

func (t TradeResp) Ack() (TradeAck, error) {
	var resp TradeAck
	if t.rawResp == nil {
		return resp, errors.New("ACK response not available")
	}
	resp = TradeAck{
		Symbol:        t.rawResp.Symbol,
		OrderId:       t.rawResp.OrderId,
		OrderListId:   t.rawResp.OrderListId,
		ClientOrderId: t.rawResp.ClientOrderId,
		TransactTime:  t.rawResp.TransactTime,
	}
	return resp, nil
}

func (t TradeResp) Result() (TradeResult, error) {
	var resp TradeResult
	if t.rawResp == nil {
		return resp, errors.New("ACK response not available")
	}
	resp = TradeResult{
		Symbol:                  t.rawResp.Symbol,
		OrderId:                 t.rawResp.OrderId,
		OrderListId:             t.rawResp.OrderListId,
		ClientOrderId:           t.rawResp.ClientOrderId,
		TransactTime:            t.rawResp.TransactTime,
		Price:                   t.rawResp.Price,
		OrigQty:                 t.rawResp.OrigQty,
		ExecutedQty:             t.rawResp.ExecutedQty,
		CummulativeQuoteQty:     t.rawResp.CummulativeQuoteQty,
		Status:                  t.rawResp.Status,
		TimeInForce:             t.rawResp.TimeInForce,
		Type:                    t.rawResp.Type,
		Side:                    t.rawResp.Side,
		WorkingTime:             t.rawResp.WorkingTime,
		SelfTradePreventionMode: t.rawResp.SelfTradePreventionMode,
	}
	return resp, nil
}

func (t TradeResp) Full() (TradeFull, error) {
	var resp TradeFull
	if t.rawResp == nil {
		return resp, errors.New("ACK response not available")
	}
	return *t.rawResp, nil
}

type TradeAck struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	OrderListId   int    `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int    `json:"transactTime"`
}

type TradeResult struct {
	Symbol                  string  `json:"symbol"`
	OrderId                 int     `json:"orderId"`
	OrderListId             int     `json:"orderListId"`
	ClientOrderId           string  `json:"clientOrderId"`
	TransactTime            int     `json:"transactTime"`
	Price                   float64 `json:"price,string"`
	OrigQty                 float64 `json:"origQty,string"`
	ExecutedQty             float64 `json:"executedQty,string"`
	CummulativeQuoteQty     float64 `json:"cummulativeQuoteQty,string"`
	Status                  string  `json:"status"`
	TimeInForce             string  `json:"timeInForce"`
	Type                    string  `json:"type"`
	Side                    string  `json:"side"`
	WorkingTime             int     `json:"workingTime"`
	SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
}

type TradeFull struct {
	Symbol                  string  `json:"symbol"`
	OrderId                 int     `json:"orderId"`
	OrderListId             int     `json:"orderListId"`
	ClientOrderId           string  `json:"clientOrderId"`
	TransactTime            int     `json:"transactTime"`
	Price                   float64 `json:"price,string"`
	OrigQty                 float64 `json:"origQty,string"`
	ExecutedQty             float64 `json:"executedQty,string"`
	CummulativeQuoteQty     float64 `json:"cummulativeQuoteQty,string"`
	Status                  string  `json:"status"`
	TimeInForce             string  `json:"timeInForce"`
	Type                    string  `json:"type"`
	Side                    string  `json:"side"`
	WorkingTime             int     `json:"workingTime"`
	SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
	Fills                   []struct {
		Price           float64 `json:"price,string"`
		Qty             float64 `json:"qty,string"`
		Commission      float64 `json:"commission,string"`
		CommissionAsset string  `json:"commissionAsset"`
		TradeId         int     `json:"tradeId"`
	} `json:"fills"`
}

func ParseResp(b []byte) (TradeFull, error) {
	var v TradeFull
	// try to unmarshal b to a TradeAck response
	err := json.Unmarshal(b, &v)
	if err != nil {
		return v, fmt.Errorf("response could not be parsed %v", string(b))
	}
	return v, nil
}
