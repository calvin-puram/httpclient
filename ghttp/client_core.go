package ghttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

func(c *client) setBody(contentType string, body interface{} ) ([]byte, error){
  if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}


func(c *client) do(method, url string, headers http.Header, body interface{}) (*http.Response, error){
	client := &http.Client{}
	fullHeaders := c.httpHeaders(headers)
	postData, err := c.setBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		log.Fatalf("unable to marshal req:%s", err.Error())
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(postData))
	if err != nil {
		log.Fatalf("unable to set req:%s", err.Error())
	}
	req.Header = fullHeaders
	return client.Do(req)
}

func(c *client) httpHeaders(headers http.Header) http.Header {
	req := make(http.Header)
  // common httpHeaders
	for ky, vl := range c.headers{
   if len(vl) > 0 {
		req.Add(ky, vl[0] )
	 }
	}
	//unique headers
	for ky, vl := range headers{
   if len(vl) > 0 {
		req.Add(ky, vl[0] )
	 }
	}
	return req
}