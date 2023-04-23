package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type CancellationResp struct {
	Symbol                  string  `json:"symbol"`
	OrigClientOrderId       string  `json:"origClientOrderId"`
	OrderId                 int     `json:"orderId"`
	OrderListId             int     `json:"orderListId"` //Unless part of an OCO, the value will always be -1.
	ClientOrderId           string  `json:"clientOrderId"`
	Price                   float64 `json:"price,string"`
	OrigQty                 float64 `json:"origQty,string"`
	ExecutedQty             float64 `json:"executedQty,string"`
	CummulativeQuoteQty     float64 `json:"cummulativeQuoteQty,string"`
	Status                  string  `json:"status"`
	TimeInForce             string  `json:"timeInForce"`
	OrderType               string  `json:"type"`
	Side                    string  `json:"side"`
	StopPrice               float64 `json:"stopPrice,string"` // might not be sent
	SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
}

type CancelOrderParams struct {
	Symbol            string
	OrderId           int
	OrigClientOrderId string
	NewClientOrderId  string
	RecvWindow        int
	Timestamp         time.Time
}

func NewCancelOrderParams(symbol string, orderId int, origClientOrderId string) CancelOrderParams {
	return CancelOrderParams{
		Symbol:            symbol,
		OrderId:           orderId,
		OrigClientOrderId: origClientOrderId,
		Timestamp:         time.Now(),
	}
}

func (c Client) CancelOrder(symbol string, orderId int) (CancellationResp, error) {
	var res CancellationResp
	query := "symbol=" + symbol +
		"&orderId=" + strconv.Itoa(orderId) +
		"&timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ORDER + "?" + query + "&signature=" + signature
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
	return res, err
}

func (c Client) CancelOrder2(symbol string, clientOrderId string) (CancellationResp, error) {
	var res CancellationResp
	query := "symbol=" + symbol +
		"&origClientOrderId=" + clientOrderId +
		"&timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ORDER + "?" + query + "&signature=" + signature
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
	return res, err
}

func (c Client) CancelOrder3(cop CancelOrderParams) (CancellationResp, error) {
	var res CancellationResp
	var oid, coid, ncoid, recvw string
	if cop.OrderId != 0 {
		oid = "&orderId=" + strconv.Itoa(cop.OrderId)
	}
	if cop.OrigClientOrderId != "" {
		coid = "&origClientOrderId=" + cop.OrigClientOrderId
	}
	if cop.NewClientOrderId != "" {
		coid = "&newClientOrderId=" + cop.NewClientOrderId
	}
	if cop.RecvWindow != 0 {
		oid = "&recvWindow=" + strconv.Itoa(cop.RecvWindow)
	}
	query := "symbol=" + cop.Symbol +
		oid + coid + ncoid + recvw +
		"&timestamp=" + strconv.Itoa(int(cop.Timestamp.UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ORDER + "?" + query + "&signature=" + signature
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
	return res, err
}
