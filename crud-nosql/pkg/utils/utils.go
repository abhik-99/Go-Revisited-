package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err == nil {
			return
		}
	}
}

func SendResponse(w http.ResponseWriter, contentType string, status int, result interface{}) {
	res, _ := json.Marshal(result)
	w.Header().Set("Content-type", contentType)
	w.WriteHeader(status)
	w.Write(res)
}
