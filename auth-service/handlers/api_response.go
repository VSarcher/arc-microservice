package handlers

import (
	"encoding/json"
	"net/http"
)

func Api_response(rw http.ResponseWriter, status int, data interface{}) {
	rw.WriteHeader(status)
	rw.Header().Add("content-type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
