package binance

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func (c Client) Trade(o Order) (TradeResp, error) {
	var res TradeResp
	trade, err := o.query(&c)
	if err != nil {
		return res, err
	}
	signature := c.Sign(trade)
	url := string(c.BaseUrl) + ORDER + "?" + trade + "&signature=" + signature
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
		e, err := ParseRespErr(b)
		if err != nil {
			return TradeResp{}, err
		}
		return res, fmt.Errorf("%v %v", e.Code, e.Msg)
	}

	tf, err := ParseResp(b)
	if err != nil {
		return res, err
	}
	res.rawResp = &tf
	return res, nil
}

func (c Client) TestTrade(o Order) (TradeResp, error) {
	var res TradeResp
	trade, err := o.query(&c)
	if err != nil {
		return res, err
	}
	signature := c.Sign(trade)
	url := string(c.BaseUrl) + TEST_ORDER + "?" + trade + "&signature=" + signature
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
		e, err := ParseRespErr(b)
		if err != nil {
			return TradeResp{}, err
		}
		return res, fmt.Errorf("%v %v", e.Code, e.Msg)
	}

	tf, err := ParseResp(b)
	if err != nil {
		return res, err
	}
	res.rawResp = &tf
	return res, nil
}
