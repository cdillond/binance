package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type OCOCancelParams struct {
	Symbol            string // REQUIRED
	OrderListId       int    // REQUIRED IF NO ListClientOrderId
	ListClientOrderId string // REQUIRED IF NO OrderListId
	NewClientOrderId  string
	RecvWindow        int
	Timestamp         time.Time
}

type OCOCancellationResp struct {
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
	OrderReports []CancellationResp `json:"orderReports"`
}

func (c Client) OCOCancelOrderList(oc OCOCancelParams) (OCOCancellationResp, error) {
	var res OCOCancellationResp
	var olid, lcoid, ncoid, recvw string
	if oc.OrderListId != 0 {
		olid = "&orderListId=" + strconv.Itoa(oc.OrderListId)
	}
	if oc.ListClientOrderId != "" {
		lcoid = "&listClientOrderId=" + oc.ListClientOrderId
	}
	if oc.NewClientOrderId != "" {
		ncoid = "&newClientOrderId=" + oc.NewClientOrderId
	}
	if oc.RecvWindow != 0 {
		recvw = "&recvWindow=" + strconv.Itoa(oc.RecvWindow)
	}
	query := "symbol=" + oc.Symbol +
		olid + lcoid + ncoid + recvw +
		"&timestamp=" + strconv.Itoa(int(oc.Timestamp.UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ORDER_LIST + "?" + query + "&signature=" + signature
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
