package handlers

import (
	"encoding/json"
	"net/http"
)

func JSON_response(rw http.ResponseWriter, status int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(data)

}
