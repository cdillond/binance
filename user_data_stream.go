package binance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type listenKey struct {
	ListenKey string `json:"listenKey`
}

// Returns a listen key and an error, which is nil on success.
func (c Client) CreateUserDataStream() (string, error) {
	var res string
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, string(c.BaseUrl)+USER_DATA, nil)
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

	var v listenKey
	err = json.Unmarshal(b, &v)
	if err != nil {
		return v.ListenKey, err
	}
	if v.ListenKey == "" {
		return v.ListenKey, errors.New("received malformed response from Binance servers")
	}
	return v.ListenKey, nil
}

// Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
func (c Client) ExtendUserDataStream(listenKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, string(c.BaseUrl)+USER_DATA+"?listenKey="+listenKey, nil)
	if err != nil {
		return err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return parseRespErr(b)
	}

	if string(b) != "{}" {
		return fmt.Errorf("received malformed response: %v", string(b))
	}
	return nil
}

func (c Client) CloseUserDataStream(listenKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, string(c.BaseUrl)+USER_DATA+"?listenKey="+listenKey, nil)
	if err != nil {
		return err
	}
	req.Header["X-MBX-APIKEY"] = []string{c.ApiKey}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// REQUEST ERROR
	if resp.StatusCode >= 400 {
		return parseRespErr(b)
	}

	if string(b) != "{}" {
		return fmt.Errorf("received malformed response: %v", string(b))
	}
	return nil
}
