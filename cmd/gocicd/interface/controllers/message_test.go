package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandlerReturnsStringTypeOfData(t *testing.T) {
	t.Run("should return: Hello Juan Perez your message will be send", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)

		data := map[string]interface{}{
			"message":       "This is a test",
			"to":            "Juan Perez",
			"from":          "Rita Asturia",
			"timeToLifeSec": 45,
		}
		dataBytes, _ := json.Marshal(data)

		c.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(dataBytes)),
		}
		Handler(c)

		want := "Hello Juan Perez your message will be send"
		got, _ := ioutil.ReadAll(res.Body)

		if string(got) != want {
			t.Errorf("got '%s' expected '%s", got, want)
		}
	})
}
