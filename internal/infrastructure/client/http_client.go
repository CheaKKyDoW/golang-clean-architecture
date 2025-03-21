package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HTTPClient struct {
	*http.Client
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   timeout,
		},
	}
}

func (c *HTTPClient) DoWithCookieToken(req *http.Request) (*http.Response, error) {
	req.AddCookie(&http.Cookie{
		// Add your cookie name and value here
	})
	return c.Client.Do(req)
}

func RequestJSONResult[T any](client *HTTPClient, method, url string, headers map[string]string, payload []byte) (statusCode int, result T, err error) {
	statusCode = http.StatusServiceUnavailable

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.DoWithCookieToken(req)
	if err != nil {
		return
	}

	if resp != nil {
		defer resp.Body.Close()
	} else {
		err = fmt.Errorf("Response is nil")
		return
	}

	statusCode = resp.StatusCode

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
