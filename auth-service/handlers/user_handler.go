package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VSarcher/arc-microservice/auth-service/config"
	"github.com/VSarcher/arc-microservice/auth-service/models"
)

// type UserHandler struct {
// }
const (
	ERR_DECODING_JSON = "Error on decoding json request"
)

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Create New User
		newUser := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			Api_response(rw, http.StatusInternalServerError, map[string]string{"error": ERR_DECODING_JSON})
			log.Fatal("Error on decoding json request")
		}

		// fmt.Println(newUser)
		config.PostgresDB.DB.Create(&newUser)

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(newUser)
	}
}
