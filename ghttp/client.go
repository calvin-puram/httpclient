package ghttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(string, interface{}, ...http.Header) (*Response, error)
	Post(string, interface{}, ...http.Header) (*Response, error)
	Patch(string, interface{}, ...http.Header) (*Response, error)
	Put(string, interface{}, ...http.Header) (*Response, error)
	Delete(string, interface{}, ...http.Header) (*Response, error)
}

type client struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *client) Get(url string, body interface{}, headers ...http.Header) (*Response, error) {

	return c.do(http.MethodGet, url, setHttpHeaders(headers...), nil)
}

func (c *client) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, setHttpHeaders(headers...), body)
}

func (c *client) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, setHttpHeaders(headers...), body)
}

func (c *client) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, setHttpHeaders(headers...), body)
}

func (c *client) Delete(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, setHttpHeaders(headers...), nil)
}
