package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/models"
)

// type UserHandler struct {
// }

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// var newUser models.User
		newUser := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&newUser); err == nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(map[string]string{"status": "error"})
			log.Fatal("Error on decoding json request")
		}

		fmt.Println(newUser)
		config.PostgresDB.DB.Create(&newUser)
		// fmt.Println(user)
		// rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(newUser)
	}
}
