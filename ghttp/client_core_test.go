package ghttp

// import (
// 	"net/http"
// 	"testing"
// )

// func TestHttpHeaders(t *testing.T) {
// 	c := clientBuilder{}
// 	commonHeaders := make(http.Header)
// 	commonHeaders.Set("Authorization", "Bearer xxx")
// 	c.SetHeaders(commonHeaders)

// 	headers := make(http.Header)
// 	headers.Set("Content-Type", "application/json")
// 	headers.Set("Accept", "application/json")

// 	fullHeaders := c.httpHeaders(headers)

// 	if len(fullHeaders) < 3 {
// 		t.Errorf("expect 3 headers but got %d headers", len(fullHeaders))
// 	}

// 	if fullHeaders.Get("Content-Type") != "application/json" {
// 		t.Errorf("expect application/json' got %q", fullHeaders.Get("Content-Type"))
// 	}

// 	if fullHeaders.Get("Accept") != "application/json" {
// 		t.Errorf("expect 'application/json' got %q", fullHeaders.Get("Accept"))
// 	}

// 	if fullHeaders.Get("Authorization") != "Bearer xxx" {
// 		t.Errorf("expect 'Bearer xxx' got %q", fullHeaders.Get("Authorization"))
// 	}
// }

// type user struct {
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// 	UserId int    `json:"userId"`
// }

// func TestSetBody(t *testing.T) {
// 	c := client{}
// 	data := user{Title: "calvin", Body: "purma", UserId: 1}
// 	t.Run("Nil Body", func(t *testing.T) {
// 		body, err := c.setBody("", nil)
// 		if err != nil {
// 			t.Errorf("didn't expect error but got %q", err.Error())
// 		}

// 		if body != nil {
// 			t.Errorf("expect nil body but got %v", string(body))
// 		}
// 	})

// 	t.Run("JSON Body", func(t *testing.T) {
// 		res := `{"title":"calvin","body":"purma","userId":1}`
// 		body, err := c.setBody("application/json", data)

// 		if err != nil {
// 			t.Errorf("didn't expect error but got %q", err.Error())
// 		}

//     if string(body) != res{
// 			t.Errorf("expect %v but got %v", res, string(body))
// 		}
// 	})

// 	t.Run("XML Body", func(t *testing.T) {
// 		res := "<user><Title>calvin</Title><Body>purma</Body><UserId>1</UserId></user>"
// 		body, err := c.setBody("application/xml", data)
// 		if err != nil {
// 			t.Errorf("didn't expect error but got %q", err.Error())
// 		}
// 		if string(body) != res{
// 			t.Errorf("expect %q but got %q", res, string(body))
// 		}
// 	})

// 	t.Run("Default Body", func(t *testing.T) {
// 		res := `{"title":"calvin","body":"purma","userId":1}`
// 		body, err := c.setBody("", data)
// 		if err != nil {
// 			t.Errorf("didn't expect error but got %q", err.Error())
// 		}

// 		if string(body) != res{
// 			t.Errorf("expect %q but got %q", res, string(body))
// 		}
// 	})
// }
