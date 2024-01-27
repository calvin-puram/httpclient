package ghttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers               http.Header
	MaxIdleConnsPerHost   int
	ResponseHeaderTimeout time.Duration
	disableTimeouts       bool
	cl                    *http.Client
	userAgent             string
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConns(mx int) ClientBuilder
	SetResponseHeaderTimeout(tm time.Duration) ClientBuilder
	SetDisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(cl *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	c := &clientBuilder{}
	return c
}

func (c *clientBuilder) Build() Client {
	return &client{
		builder: c,
	}
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {

	c.headers = headers
	return c
}

func (c *clientBuilder) SetMaxIdleConns(mx int) ClientBuilder {
	c.MaxIdleConnsPerHost = mx
	return c
}

func (c *clientBuilder) SetResponseHeaderTimeout(tm time.Duration) ClientBuilder {
	c.ResponseHeaderTimeout = tm
	return c
}

func (c *clientBuilder) SetDisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(cl *http.Client) ClientBuilder {
	if cl != nil {
		c.cl = cl
	}
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	if userAgent != "" {
		c.userAgent = userAgent
	}
	return c
}
