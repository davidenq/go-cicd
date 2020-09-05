package end2end

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

type DataRequest struct {
	Method             string
	Endpoint           string
	Data               map[string]interface{}
	ExpectedStatusCode int
	ExpectedMessage    string
	CheckStatusCode    bool
	t                  *testing.T
	APIKey             string
}

var host = os.Getenv("HOST")
var devopsEndpoint = host + "/DevOps"
var healthEndpoint = host + "/health"

func TestHTTPSTatusCodeForNonAuthenticatedRequests(t *testing.T) {
	t.Run("should return HTTP Status Code 401 for non authenticated GET request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "GET",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusUnauthorized,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             "",
		}
		assetRequest(dataRequest)
	})

	t.Run("should return HTTP Status Code 401 for non authenticated POST request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "POST",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusUnauthorized,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             "",
		}
		assetRequest(dataRequest)
	})

	/*t.Run("should return HTTP Status Code 401 for non authenticated HEAD request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "HEAD",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusUnauthorized,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             "",
		}
		assetRequest(dataRequest)
	})*/

	t.Run("should return HTTP Status Code 401 for non authenticated PUT request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "PUT",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusUnauthorized,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             "",
		}
		assetRequest(dataRequest)

	})
	t.Run("should return HTTP Status Code 401 for non authenticated DELETE request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "DELETE",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusUnauthorized,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             "",
		}
		assetRequest(dataRequest)
	})
}

const apikey = "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c"

func TestHTTPSTatusCodeForAuthenticatedRequests(t *testing.T) {
	t.Run("should return HTTP Status Code 405 (Method Not Allowed) for authenticated GET request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "GET",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	t.Run("should return HTTP Status Code 200 for authenticated POST request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "POST",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusOK,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	/*t.Run("should return HTTP Status Code 405 (Method Not Allowed) for authenticated HEAD request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "HEAD",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})*/
	t.Run("should return HTTP Status Code 405 (Method Not Allowed) for authenticated PUT request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "PUT",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	t.Run("should return HTTP Status Code 405 (Method Not Allowed) for authenticated DELETE request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "DELETE",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
}

func TestStringResponseForAuthenticatedRequest(t *testing.T) {
	t.Run("should response 'ERROR' message for GET request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "GET",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "ERROR",
			CheckStatusCode:    false,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	t.Run("should response 'Hello Juan Perez your message will be send' message", func(t *testing.T) {
		data := map[string]interface{}{
			"message":       "This is a test",
			"to":            "Juan Perez",
			"from":          "Rita Asturia",
			"timeToLifeSec": 45,
		}
		dataRequest := &DataRequest{
			Method:             "POST",
			Endpoint:           devopsEndpoint,
			Data:               data,
			ExpectedStatusCode: http.StatusOK,
			ExpectedMessage:    "Hello Juan Perez your message will be send",
			CheckStatusCode:    false,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	/*t.Run("should response 'ERROR' message for HEAD request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "HEAD",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "ERROR",
			CheckStatusCode:    false,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	*/
	t.Run("should response 'ERROR' message for PUT request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "PUT",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "ERROR",
			CheckStatusCode:    false,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
	t.Run("should response 'ERROR' message for DELETE request", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "DELETE",
			Endpoint:           devopsEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusMethodNotAllowed,
			ExpectedMessage:    "ERROR",
			CheckStatusCode:    false,
			t:                  t,
			APIKey:             apikey,
		}
		assetRequest(dataRequest)
	})
}

func TestHealthEndpoint(t *testing.T) {
	t.Run("should return http code 200 to check /health endpoint", func(t *testing.T) {
		dataRequest := &DataRequest{
			Method:             "GET",
			Endpoint:           healthEndpoint,
			Data:               nil,
			ExpectedStatusCode: http.StatusOK,
			ExpectedMessage:    "",
			CheckStatusCode:    true,
			t:                  t,
		}
		assetRequest(dataRequest)
	})
}

func assetRequest(dr *DataRequest) {
	dataJSON, _ := json.Marshal(dr.Data)
	data := bytes.NewBuffer(dataJSON)

	client := http.Client{}
	request, _ := http.NewRequest(dr.Method, dr.Endpoint, data)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Parse-REST-API-Key", dr.APIKey)

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	if dr.CheckStatusCode {
		if response.StatusCode != dr.ExpectedStatusCode {
			dr.t.Errorf("got '%d' expected '%d", response.StatusCode, dr.ExpectedStatusCode)
		}
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		stringBody := string(body)
		if stringBody != dr.ExpectedMessage {
			dr.t.Errorf("got '%s' expected '%s", stringBody, dr.ExpectedMessage)
		}
	}
}
