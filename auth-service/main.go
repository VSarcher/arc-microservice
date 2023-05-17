package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hi")
	// })
	// log.Fatal(http.ListenAndServe(":8081", nil))

	Router := mux.NewRouter()
	Router.HandleFunc("/", HomeHandler()).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", Router))
}

func HomeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		aa := `{"apple": "app"}`
		json.NewEncoder(rw).Encode(aa)
	}
}