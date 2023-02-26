package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type UserTradeResp struct {
	Symbol          string  `json:"symbol"`
	Id              int     `json:"id"`
	OrderId         int     `json:"orderId"`
	OrderListId     int     `json:"orderListId"`
	Price           float64 `json:"price,string"`
	Qty             float64 `json:"qty,string"`
	QuoteQty        float64 `json:"quoteQty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
	Time            int     `json:"time"`
	IsBuyer         bool    `json:"isBuyer"`
	IsMaker         bool    `json:"isMaker"`
	IsBestMatch     bool    `json:"isBestMatch"`
}

type UserTradesParams struct {
	Symbol     string
	OrderId    int
	StartTime  time.Time
	EndTime    time.Time
	FromId     int
	Limit      int
	RecvWindow int
	Timestamp  time.Time
}

func NewUserTradesParams(symbol string) UserTradesParams {
	return UserTradesParams{
		Symbol:    symbol,
		Timestamp: time.Now(),
	}
}

func queryString(u UserTradesParams) string {
	res := "symbol=" + u.Symbol +
		"&timestamp=" + strconv.Itoa(int(u.Timestamp.UnixMilli()))
	if u.Limit != 0 {
		res += "&limit=" + strconv.Itoa(u.Limit)
	}
	if u.RecvWindow != 0 {
		res += "&recvWindow=" + strconv.Itoa(u.RecvWindow)
	}

	if u.OrderId != 0 {
		res += "&orderId=" + strconv.Itoa(u.OrderId)
	}
	if u.FromId != 0 {
		res += "&fromId" + strconv.Itoa(u.FromId)
		return res
	}

	if !u.StartTime.IsZero() {
		res += "&startTime=" + strconv.Itoa(int(u.StartTime.UnixMilli()))
	}
	if !u.EndTime.IsZero() {
		res += "&endTime=" + strconv.Itoa(int(u.EndTime.UnixMilli()))
	}
	return res
}

func (c Client) UserTrades(u UserTradesParams) ([]UserTradeResp, error) {
	var res []UserTradeResp
	query := queryString(u)
	signature := c.Sign(query)
	url := string(c.BaseUrl) + USER_TRADES + "?" + query + "&signature=" + signature
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
		return res, parseRespErr(b)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}

	return res, nil

}
