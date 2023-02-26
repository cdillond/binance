package binance

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type ExchangeInfoResp struct {
	Timezone           string      `json:"timezone"`
	ServerTime         int64       `json:"serverTime"`
	RateLimits         []RateLimit `json:"rateLimits"`
	RawExchangeFilters []any       `json:"exchangeFilters"`
	//ExchangeFilters    map[string]any
	Symbols            []Symbol `json:"symbols"`
	Permissions        []string `json:"permissions"`
	DefaultSeltPrevent string   `json:"defaultSelfTradePreventionMode"`
	AllowedSelfPrevent []string `json:"allowedSelfTradePreventionModes"`
}

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

func (c Client) ExchangeInfo() (ExchangeInfoResp, error) {
	return c.exchangeInfo("")
}

func (c Client) SymbolInfo(symbol string) (ExchangeInfoResp, error) {
	return c.exchangeInfo("?symbol=" + symbol)
}

func (c Client) SymbolsInfo(symbols []string) (ExchangeInfoResp, error) {
	b, err := json.Marshal(symbols)
	if err != nil {
		return ExchangeInfoResp{}, err
	}
	return c.exchangeInfo("?symbols=" + string(b))
}

func (c Client) PermissionsInfo(permission string) (ExchangeInfoResp, error) {
	return c.exchangeInfo("?permissions=" + permission)
}

// The permissions parameter is essentially useless for Binance.US
func (c Client) exchangeInfo(queryString string) (ExchangeInfoResp, error) {
	url := string(c.BaseUrl) + EXCHANGE_INFO + queryString

	var eir ExchangeInfoResp
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return eir, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return eir, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return eir, err
	}
	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return eir, parseRespErr(b)
	}

	err = json.Unmarshal(b, &eir)
	//fmt.Println(eir, err)
	return eir, err
}
