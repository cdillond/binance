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

type Kline struct {
	OpenTime         time.Time
	Open             float64
	High             float64
	Low              float64
	Close            float64
	Volume           float64
	CloseTime        time.Time
	QuoteAssetVol    float64
	NumTrades        int
	TakerBuyBaseVol  float64
	TakerBuyQuoteVol float64
}

// The fields for most return values are of an equivalent type as the original JSON response from the API;
// not so with Klines. Because of the way the response is structured, it makes sense to typecast the
// unmarshalled JSON to numeric types that are easier to work with.
func (c *Client) Klines(symbol, interval string, limit int, startTime, endTime time.Time) ([]Kline, error) {
	s, e := startTime.UnixMilli(), endTime.UnixMilli()
	q := "?symbol=" + symbol +
		"&interval=" + interval +
		"&limit=" + strconv.Itoa(limit) +
		"&startTime=" + strconv.Itoa(int(s)) +
		"&endTime=" + strconv.Itoa(int(e))
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+KLINES+q, nil)
	if err != nil {
		return []Kline{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []Kline{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Kline{}, err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		e, err := ParseRespErr(b)
		if err != nil {
			return []Kline{}, err
		}
		return []Kline{}, fmt.Errorf("%v %v", e.Code, e.Msg)
	}
	var tr [][]any
	err = json.Unmarshal(b, &tr)
	klines := make([]Kline, len(tr))
	for i := 0; i < len(tr); i++ {
		k, err := unmarshalKline(tr[i])
		if err != nil {
			return []Kline{}, err
		}
		klines[i] = k
	}
	return klines, nil
}

// VERY TEDIOUS
func unmarshalKline(raw []any) (Kline, error) {
	var res Kline
	if len(raw) < 11 {
		return res, fmt.Errorf("malformed raw kline response")
	}
	openTimeF, ok := raw[0].(float64)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}

	openS, ok := raw[1].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	openF, err := strconv.ParseFloat(openS, 64)
	if err != nil {
		return res, err
	}

	highS, ok := raw[2].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	highF, err := strconv.ParseFloat(highS, 64)
	if err != nil {
		return res, err
	}

	lowS, ok := raw[3].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	lowF, err := strconv.ParseFloat(lowS, 64)
	if err != nil {
		return res, err
	}

	closeS, ok := raw[4].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	closeF, err := strconv.ParseFloat(closeS, 64)
	if err != nil {
		return res, err
	}

	volumeS, ok := raw[5].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	volumeF, err := strconv.ParseFloat(volumeS, 64)
	if err != nil {
		return res, err
	}

	closeTimeF, ok := raw[6].(float64)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}

	qavS, ok := raw[7].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	qavF, err := strconv.ParseFloat(qavS, 64)
	if err != nil {
		return res, err
	}

	numTradesF, ok := raw[8].(float64)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}

	tbbavS, ok := raw[9].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	tbbavF, err := strconv.ParseFloat(tbbavS, 64)
	if err != nil {
		return res, err
	}

	tbbqvS, ok := raw[10].(string)
	if !ok {
		return res, fmt.Errorf("unable to parse kline response")
	}
	tbbqvF, err := strconv.ParseFloat(tbbqvS, 64)
	if err != nil {
		return res, err
	}

	return Kline{
		OpenTime:         time.Unix(0, int64(openTimeF)*1_000_000),
		Open:             openF,
		High:             highF,
		Low:              lowF,
		Close:            closeF,
		Volume:           volumeF,
		CloseTime:        time.Unix(0, int64(closeTimeF)*1_000_000),
		QuoteAssetVol:    qavF,
		NumTrades:        int(numTradesF),
		TakerBuyBaseVol:  tbbavF,
		TakerBuyQuoteVol: tbbqvF,
	}, nil
}
