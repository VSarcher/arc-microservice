package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/models"
)

// type UserHandler struct {
// }

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Create New User
		newUser := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(map[string]string{"status": "error"})
			log.Fatal("Error on decoding json request")
		}

		// fmt.Println(newUser)
		config.PostgresDB.DB.Create(&newUser)

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(newUser)
	}
}
