package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/davidenq/go-cicd/cmd/gocicd/types"
)

func TestHandlerReturnsStringTypeOfData(t *testing.T) {
	t.Run("should return: Hello Juan Perez your message will be send", func(t *testing.T) {
		res := httptest.NewRecorder()
		data := &types.BodyMessage{
			Message:       "This is a test",
			To:            "Juan Perez",
			From:          "Rita Asturia",
			TimeToLifeSec: 45,
		}
		dataBytes, _ := json.Marshal(data)
		req, _ := http.NewRequest("POST", "http://localhost:4000/DevOps", strings.NewReader(string(dataBytes)))
		Handler(res, req)
		var want string
		got, _ := ioutil.ReadAll(res.Body)
		if reflect.TypeOf(string(got)) != reflect.TypeOf(want) {
			t.Errorf("got '%s' expected '%s", got, want)
		}
	})
}
