package examples

import (
	"net/http"

	"github.com/calvin-puram/httpclient/ghttp"
)

var (
	httpClient = getClient()
)

func getClient() ghttp.Client {
	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "application/json")


	client := ghttp.NewBuilder().SetHttpClient(&http.Client{}).SetHeaders(commonHeader).SetUserAgent("calvintest").Build()
	return client
}
