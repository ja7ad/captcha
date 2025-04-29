package captcha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	httpClient *http.Client
}

func newClient() *client {
	return &client{
		httpClient: &http.Client{
			Transport: baseTransport(),
			Timeout:   10 * time.Second,
		},
	}
}

func (c *client) Request(
	ctx context.Context,
	endpoint, method string,
	body any,
	resp any,
	headers map[string]string,
	params map[string]string,
) error {
	var reqBody io.Reader

	if body != nil {
		switch v := body.(type) {
		case url.Values:
			reqBody = bytes.NewBufferString(v.Encode())
			// if Content-Type not already set, set it
			if headers == nil {
				headers = map[string]string{}
			}
			if _, ok := headers["Content-Type"]; !ok {
				headers["Content-Type"] = "application/x-www-form-urlencoded"
			}
		default:
			data, err := json.Marshal(body)
			if err != nil {
				return err
			}
			reqBody = bytes.NewBuffer(data)

			if headers == nil {
				headers = map[string]string{}
			}
			if _, ok := headers["Content-Type"]; !ok {
				headers["Content-Type"] = "application/json"
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, reqBody)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if len(params) > 0 {
		query := url.Values{}
		for k, v := range params {
			query.Set(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New("non-2xx response: " + res.Status)
	}

	if resp != nil {
		return json.NewDecoder(res.Body).Decode(resp)
	}

	return nil
}

func baseTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
