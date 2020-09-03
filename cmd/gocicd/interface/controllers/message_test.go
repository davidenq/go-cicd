package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandlerReturnsStringTypeOfData(t *testing.T) {
	t.Run("should return a string type of data when receives a specific json data structure", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "http://localhost:4000/DevOps", nil)
		Handler(res, req)

		var want string
		got, _ := ioutil.ReadAll(res.Body)
		if reflect.TypeOf(got) != reflect.TypeOf(want) {
			t.Errorf("got '%s' expected '%s", got, want)
		}
	})
}
