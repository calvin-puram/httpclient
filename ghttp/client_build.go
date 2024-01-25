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
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConns(mx int) ClientBuilder
	SetResponseHeaderTimeout(tm time.Duration) ClientBuilder
	SetDisableTimeouts(disable bool) ClientBuilder
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
