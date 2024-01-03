package ghttp

import "net/http"

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header, body interface{}) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header, body interface{}) (*http.Response, error)
}

type client struct{
	headers http.Header
}

func New() HttpClient{
	c := &client{}
	return c
}

func (c *client) SetHeaders(headers http.Header) {
  c.headers = headers
}

func (c *client) Get(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}


func (c *client) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *client) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *client) Delete(url string, headers http.Header, body interface{}) (*http.Response, error) {
  return c.do(http.MethodDelete, url, headers, nil)
}