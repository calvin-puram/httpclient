package examples

import "fmt"

type githubRes struct{
	Login string 	`json:"login"`
	URL string `json:"url"`
	Bio string `json:"bio,omitempty"`
}

func GetGithubUser() (*string, error) {
	 res, err := httpClient.Get("https://api.github.com/user", nil, nil)

	 if err != nil {
		return nil, err
	 }

	 fmt.Println("response status", res.Status())
	 fmt.Println("response statuscode", res.StatusCode())
	 fmt.Println("response body", res.BodyString())
	 var userEndpoint githubRes
	 if err := res.UnmarsalJSON(&userEndpoint); err != nil {
		return nil, err
	 }

	 return &userEndpoint.URL, nil
}