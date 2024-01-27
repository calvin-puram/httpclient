package ghttp

import (
	"net/http"

	"github.com/calvin-puram/httpclient/config"
)

func setHttpHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
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

	if req.Get(config.HeaderUserAgent) == "" {
		req.Set(config.HeaderUserAgent, c.builder.userAgent)
	}

	return req
}
