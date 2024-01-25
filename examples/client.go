package examples

import (
	"net/http"
	"time"

	"github.com/calvin-puram/httpclient/ghttp"
)

var (
	httpClient = getClient()
)

func getClient() ghttp.Client {
	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "application/json")
	commonHeader.Set("Authorization", "Bearer "+ "xxx")

	client := ghttp.NewBuilder().SetDisableTimeouts(false).SetMaxIdleConns(5).SetResponseHeaderTimeout(3*time.Second).SetHeaders(commonHeader).Build()
	return client
}
