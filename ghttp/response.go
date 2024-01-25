package ghttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	header     http.Header
	body       []byte
}

func (r Response) Status() string {
	return r.status
}

func (r Response) StatusCode() int{
	return r.statusCode
}

func (r Response) Header() http.Header{
	return r.header
}

func (r Response) Bytes() []byte {
	return r.body
}

func (r Response) BodyString() string{
	return string(r.body)
}

func (r Response) UnmarsalJSON(target interface{}) error {
  return json.Unmarshal(r.body, target)
}
