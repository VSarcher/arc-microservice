package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"auth-service/config"
	"auth-service/handlers"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

func main() {
	rand.Intn(10)
	fmt.Println(os.Getenv("DB_USER"))
	
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Error in connecting with DB: %s\n", err.Error())
	}
	Router := mux.NewRouter()
	Router.HandleFunc("/", HomeHandler()).Methods("GET")

	Router.HandleFunc("/api/user/signup", handlers.CreateUser()).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", Router))
}

func HomeHandler() http.HandlerFunc {
	// a := models.User{}
	// fmt.Println(a)
	return func(rw http.ResponseWriter, r *http.Request) {
		aa := `{"apple": "app"}`
		json.NewEncoder(rw).Encode(aa)
	}
}
