package ghttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers http.Header, body interface{}) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header, body interface{}) (*Response, error)
}

type client struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *client) Get(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *client) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *client) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *client) Delete(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
