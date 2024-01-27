package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       []byte
}

func (r Response) Bytes() []byte {
	return r.Body
}

func (r Response) BodyString() string {
	return string(r.Body)
}

func (r Response) UnmarsalJSON(target interface{}) error {
	return json.Unmarshal(r.Body, target)
}
