package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/calvin-puram/httpclient/config"
	"github.com/calvin-puram/httpclient/config/mocks"
)

func TestCreateTodo(t *testing.T) {

	t.Run("POST Todo check response error", func(t *testing.T) {
		mocks.FlushMock()

		mocks.AddMock(mocks.Mock{
			URL:         config.BaseURL,
			Method:      http.MethodPost,
			RequestBody: `{"id":1,"title":"test todo","body":"clean code","userId":1}`,
			Error:       errors.New("post request timeout"),
		})

		body, err := createTodo()

		if body != nil {
			t.Errorf("didn't expect a body but got %v", body)
		}

		if err == nil {
			t.Error("expect an error but got nil")
		}

		if err.Error() != "post request timeout" {
			t.Errorf("expect `post request timeout` but got %q", err.Error())
		}

	})

	t.Run("POST Todo response body", func(t *testing.T) {
		mocks.FlushMock()

		mocks.AddMock(mocks.Mock{
			Method:             http.MethodPost,
			URL:                config.BaseURL,
			RequestBody:        `{"id":1,"title":"test todo","body":"clean code","userId":1}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":1,"title":"test todo","body":"clean code","userId":1}`,
		})

		body, err := createTodo()

		if err != nil {
			t.Errorf("didn't expect an error but go %v", err.Error())
		}

		if body == nil {
			t.Error("expected an body but got nil")
		}

	})

}
