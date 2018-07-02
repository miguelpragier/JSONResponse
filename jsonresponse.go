package jsonresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse is a simple structure designed for support marshalling targeting http api response.
type JSONResponse struct {
	// If this api operation was sucessful
	Success bool `json:"success"`
	// the expected payload
	Data interface{} `json:"data"`
	// optional message, if something goes wrong
	Message string `json:"message"`
	// result status code
	Code int `json:"code"`
}

var lastJSON = ""

// Answer formats and returns a well formed json containing a success code, a string-map, a message and a code.
func Answer(w http.ResponseWriter, success bool, data interface{}, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	jr := JSONResponse{success, data, message, code}

	if jb, err := json.Marshal(jr); err == nil {
		lastJSON = string(jb)

		w.WriteHeader(code)
		w.Write(jb)
	} else {
		w.WriteHeader(http.StatusInternalServerError)

		j := `{"success":false,"data":nil,"message":"%s","code":500}`

		lastJSON = fmt.Sprintf(j, err.Error())

		w.Write([]byte(lastJSON))
	}
}

// LastJSON returns the last generated json string
func LastJSON() string {
	return lastJSON
}
