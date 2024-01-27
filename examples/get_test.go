package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/calvin-puram/httpclient/config"
	"github.com/calvin-puram/httpclient/ghttp"
)

func TestMain(m *testing.M) {
	ghttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetTodo(t *testing.T) {

	t.Run("GET Todo return an error", func(t *testing.T) {
		ghttp.FlushMock()

		ghttp.AddMock(ghttp.Mock{
			URL:    fmt.Sprintf("%s/%d", config.BaseURL, 1),
			Method: http.MethodGet,
			Error:  errors.New("request timeout"),
		})

		body, err := getTodoById()

		if body != nil {
			t.Errorf("didn't expect a body but got %v", body)
		}

		if err == nil {
			t.Error("expect an error but got nil")
		}

		if err.Error() != "request timeout" {
			t.Errorf("expect `request timeout` but got %q", err.Error())
		}
	})

	t.Run("GET Todo marshal body error", func(t *testing.T) {
		ghttp.FlushMock()

		ghttp.AddMock(ghttp.Mock{
			Method:             http.MethodGet,
			URL:                fmt.Sprintf("%s/%d", config.BaseURL, 1),
			ResponseBody:       `{"id":1,"title":"test","body": 34}`,
			ResponseStatusCode: http.StatusOK,
		})

		body, err := getTodoById()

		if body != nil {
			t.Errorf("didn't expect a body but got %v", body)
		}

		if err == nil {
			t.Error("expect an error but got nil")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct") {
			t.Errorf("expect `json unmarshal error` but got %v", err.Error())
		}
	})

	t.Run("GET Todo response body", func(t *testing.T) {
		ghttp.FlushMock()

		ghttp.AddMock(ghttp.Mock{
			Method:             http.MethodGet,
			URL:                fmt.Sprintf("%s/%d", config.BaseURL, 1),
			ResponseBody:       `{"id":1,"title":"test","body":"todo body"}`,
			ResponseStatusCode: http.StatusOK,
		})

		body, err := getTodoById()

		if err != nil {
			t.Errorf("didn't expect an error but go %v", err.Error())
		}

		if body == nil {
			t.Error("expected an body but got nil")
		}

	})
}
