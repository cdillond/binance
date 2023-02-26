package binance

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//func ParseRespErr(b []byte) (RespErr, error) {
//	var e RespErr
//	err := json.Unmarshal(b, &e)
//	return e, err
//}

func parseRespErr(b []byte) error {
	var e respErr
	err := json.Unmarshal(b, &e)
	if err != nil {
		return fmt.Errorf("%w %v", ErrRespParse, string(b))
	}
	err = codeToError(e.Code)
	return fmt.Errorf("%w %v", err, e.Msg)

}

func (c Client) Sign(s string) string {
	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	mac.Write([]byte(s))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// Returns a nil error on successful connection.
func (c Client) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.RequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(c.BaseUrl)+PING, nil)
	if err != nil {
		return err
	}
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
	return nil
}

// Binance's API uses only milliseconds since the Unix Epoch.
// This function converts ints representing Binance's time format to time.Time objects.
func Itot(i int) time.Time {
	return time.Unix(0, int64(i)*1_000_000)
}
