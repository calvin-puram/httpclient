package ghttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/calvin-puram/httpclient/config"
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
	case config.ContentTypeJSON:
		return json.Marshal(body)
	case config.ContentTypeXML:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *client) do(method, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.httpHeaders(headers)

	postData, err := c.setBody(fullHeaders.Get(config.HeaderContentType), body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(postData))
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
		status:     res.Status,
		statusCode: res.StatusCode,
		header:     res.Header,
		body:       data,
	}, nil

}

func (c *client) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.cl != nil {
			c.client = c.builder.cl
			return
		}

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
