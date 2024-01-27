package mocks

import (
	"fmt"
	"net/http"

	"github.com/calvin-puram/httpclient/config/core"
)

type Mock struct {
	URL         string
	RequestBody string
	Method      string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

func (m *Mock) GetResBody() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	return &core.Response{
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
		Status:     fmt.Sprintf("%v %v", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
	}, nil
}
