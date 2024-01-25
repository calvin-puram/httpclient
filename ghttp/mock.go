package ghttp

import (
	"fmt"
	"net/http"
)


type Mock struct{
	URL string
	RequestBody string
	Method string

	Error error
	ResponseBody string
	ResponseStatusCode int
}

func (m *Mock) GetResBody() (*Response, error){
	if m.Error != nil {
		return nil, m.Error
	}

	return &Response{
		statusCode: m.ResponseStatusCode,
		body: []byte(m.ResponseBody),
		status: fmt.Sprintf("%v %v", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
	}, nil
}

