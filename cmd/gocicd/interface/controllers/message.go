package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidenq/go-cicd/cmd/gocicd/domain"
	"github.com/davidenq/go-cicd/cmd/gocicd/types"
)

//Handler .
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	data, _ := ioutil.ReadAll(r.Body)
	dataStruct := &types.BodyMessage{}
	json.Unmarshal(data, &dataStruct)
	message := domain.GenerateMessage(dataStruct)
	fmt.Fprintf(w, message)
}
