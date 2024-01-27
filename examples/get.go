package examples

import (
	"fmt"

	"github.com/calvin-puram/httpclient/config"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body,omitempty"`
}

func getTodoById() (*Todo, error) {
	res, err := httpClient.Get(fmt.Sprintf("%s/%d", config.BaseURL, 1), nil)

	if err != nil {
		return nil, err
	}

	var todo Todo
	if err := res.UnmarsalJSON(&todo); err != nil {
		return nil, err
	}

	return &todo, nil
}
