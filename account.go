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

type Account struct {
	MakerCommission  int `json:"makerCommission"`
	TakerCommission  int `json:"takerCommission"`
	BuyerCommission  int `json:"buyerCommission"`
	SellerCommission int `json:"sellerCommission"`
	CommissionRates  struct {
		Maker  float64 `json:"maker,string"`
		Taker  float64 `json:"taker,string"`
		Buyer  float64 `json:"buyer,string"`
		Seller float64 `json:"seller,string"`
	} `json:"commissionRates"`
	CanTrade                bool   `json:"canTrade"`
	CanWithdraw             bool   `json:"canWithdraw"`
	CanDeposit              bool   `json:"canDeposit"`
	Brokered                bool   `json:"brokered"`
	RequireSelfTradePrevent bool   `json:"requireSelfTradePrevention"`
	UpdateTime              int    `json:"updateTime"`
	AccountType             string `json:"accountType"`
	Balances                []struct {
		Asset  string  `json:"asset"`
		Free   float64 `json:"free,string"`
		Locked float64 `json:"locked,string"`
	} `json:"balances"`
	Permissions []string `json:"permissions"`
}

func (c Client) Account() (Account, error) {
	// does not allow for recvWindow param
	var res Account
	query := "timestamp=" + strconv.Itoa(int(time.Now().UnixMilli()))
	signature := c.Sign(query)
	url := string(c.BaseUrl) + ACCT + "?" + query + "&signature=" + signature
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
