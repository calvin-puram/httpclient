package examples

import (
	"fmt"

	"github.com/calvin-puram/httpclient/config"
)

type resTodo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body,omitempty"`
	UserId int    `json:"userId"`
}

var (
	reqBody = resTodo{
		Id:     1,
		Title:  "test todo",
		Body:   "clean code",
		UserId: 1,
	}
)

func createTodo() (*string, error) {
	res, err := httpClient.Post(config.BaseURL, reqBody)

	if err != nil {
		return nil, err
	}

	var resBody Todo
	if err := res.UnmarsalJSON(&resBody); err != nil {
		return nil, fmt.Errorf("post unmarshal err: %v", err.Error())
	}

	return &resBody.Title, nil
}
