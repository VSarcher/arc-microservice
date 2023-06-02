package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VSarcher/arc-microservice/auth-service/config"
	"github.com/VSarcher/arc-microservice/auth-service/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file for configuration.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect Postgres
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Error in connecting with DB: %s\n", err.Error())
	}
	// Config Router.
	Router := mux.NewRouter()

	Router.HandleFunc("/", HomeHandler()).Methods("GET")
	Router.HandleFunc("/api/user/signup", handlers.CreateUser()).Methods("POST")
	Router.HandleFunc("/api/user/signin", handlers.LoginUser()).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", Router))
}

func HomeHandler() http.HandlerFunc {
	// a := models.User{}
	// fmt.Println(a)
	return func(rw http.ResponseWriter, r *http.Request) {
		aa := `{"apple": "app"}`
		jsonResponse, err := json.Marshal(aa)
		if err != nil {
			log.Fatalf(err.Error())
		}
		json.NewEncoder(rw).Encode(jsonResponse)
	}
}
