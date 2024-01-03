package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/calvin-puram/httpclient/ghttp"
)

type user struct {
	Title string `json:"title"`
	Body  string    `json:"body"`
	UserId int `json:"userId"`
}

func githubClient() ghttp.HttpClient {
	c := ghttp.New()

	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "application/json")
	commonHeader.Set("Accept", "application/json")
	// commonHeader.Set("Authorization", "Bearer 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")
	c.SetHeaders(commonHeader)

	return c

}

func main() {
	c := githubClient()
	//get
	// getgithub(c)

	//Post
	createGithub(c)

}

func getgithub(c ghttp.HttpClient) {
	res, err := c.Get("https://api.github.com/user", nil, nil)
	if err != nil {
		log.Fatalf("unable to get resp: %s", err.Error())
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}

func createGithub(c ghttp.HttpClient) {
	resp, err := c.Post("https://jsonplaceholder.typicode.com/posts", nil, user{Title: "calvin", Body: "xxy", UserId: 1})
	if err != nil {
		log.Fatalf("unable to get resp: %s", err.Error())
	}
	defer resp.Body.Close()
  data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
