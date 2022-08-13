package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, contentType string, status int, result interface{}) {
	res, _ := json.Marshal(result)
	w.Header().Set("Content-type", contentType)
	w.WriteHeader(status)
	w.Write(res)
}
