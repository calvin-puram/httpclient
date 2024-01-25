package ghttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConns          = 5
	defaultResponseHeaderTimeout = 3 * time.Second
)

func (c *client) setBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *client) do(method, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.httpHeaders(headers)
	postData, err := c.setBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	if mock := mockedserver.getMock(method, url, string(postData)); mock != nil {
    return mock.GetResBody()
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header = fullHeaders
	client := c.getHttpClient()
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return &Response{
		status: res.Status,
		statusCode: res.StatusCode,
		header: res.Header,
		body: data,
	}, nil


}

func (c *client) httpHeaders(headers http.Header) http.Header {
	req := make(http.Header)
	// common httpHeaders
	for ky, vl := range c.builder.headers {
		if len(vl) > 0 {
			req.Add(ky, vl[0])
		}
	}
	//unique headers
	for ky, vl := range headers {
		if len(vl) > 0 {
			req.Add(ky, vl[0])
		}
	}
	return req
}

func (c *client) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.builder.ResponseHeaderTimeout,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnsPerHost(),
				ResponseHeaderTimeout: c.getResponseHeaderTimeout(),
			},
		}
	})

	return c.client
}

func (c *client) getMaxIdleConnsPerHost() int {
	if c.builder.MaxIdleConnsPerHost > 0 {
		return c.builder.MaxIdleConnsPerHost
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultMaxIdleConns
}

func (c *client) getResponseHeaderTimeout() time.Duration {
	if c.builder.ResponseHeaderTimeout > 0 {
		return c.builder.ResponseHeaderTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseHeaderTimeout
}
