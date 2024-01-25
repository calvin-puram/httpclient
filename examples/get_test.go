package examples

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/calvin-puram/httpclient/ghttp"
)

func TestMain(m *testing.M) {
	ghttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetGithubEndpoint(t *testing.T) {
	t.Run("github enpoint return err", func(t *testing.T) {
		ghttp.FlushMock()
		//initialization
		ghttp.AddMock(ghttp.Mock{
			URL: "https://api.github.com/user",
			Method: http.MethodGet,
			Error: errors.New("request timeout"),
		})

		//execution
		endpoint, err := GetGithubUser()

		//validation
		if endpoint != nil {
			t.Errorf("didn't expect an endpoint but got %v", endpoint)
		}

		if err == nil {
			t.Error("expect an error but got nil")
		}

		if err.Error() != "request timeout"{
			t.Errorf("expect `request timeout` but got %q", err.Error())
		}
	})

	t.Run("marshal body err", func(t *testing.T) {
		ghttp.FlushMock()
		//initialization
		ghttp.AddMock(ghttp.Mock{
			Method: http.MethodGet,
			URL: "https://api.github.com/user",
			ResponseBody: `{"url":123}`,
      ResponseStatusCode: http.StatusOK,
		})

		//execution
		endpoint, err := GetGithubUser()

		//validation
		if endpoint != nil {
			t.Errorf("didn't expect an endpoint but got %v", endpoint)
		}

		if err == nil {
			t.Error("expect an error but got nil")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct") {
			t.Errorf("expect `json unmarshal error` but got %v", err.Error())
		}
	})

	t.Run("response endpoint", func(t *testing.T) {
		ghttp.FlushMock()
		//initialization
		ghttp.AddMock(ghttp.Mock{
			Method: http.MethodGet,
			URL: "https://api.github.com/user",
			ResponseBody: `{"url":"https://api.github.com/users/blockops-engineering"}`,
      ResponseStatusCode: http.StatusOK,
		})

		//execution
		endpoint, err := GetGithubUser()

		//validation
		if err != nil {
			t.Errorf("didn't expect an error but go %v", err)
		}

		if endpoint == nil {
			t.Error("expected an endpoint but got nil")
		}

		if endpoint == nil {
			t.Error("expected an endpoint but got nil")
		}

	})
}