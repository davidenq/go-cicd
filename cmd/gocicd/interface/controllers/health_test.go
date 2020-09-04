package controllers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthHandler(t *testing.T) {

	t.Run("should return a HTTP Status Code 200", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		Health(c)

		want := 200
		got := res.Code
		if got != want {
			t.Errorf("got '%d' expected '%d", got, want)
		}
	})

	t.Run("should return a string: ok", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		Health(c)

		want := "ok"
		got, _ := ioutil.ReadAll(res.Body)

		if string(got) != want {
			t.Errorf("got '%s' expected '%s", got, want)
		}
	})

}
