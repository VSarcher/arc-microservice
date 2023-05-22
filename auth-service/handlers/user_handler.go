package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/models"
)

// type UserHandler struct {
// }

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// print
		var newUser models.User
		json.NewDecoder(r.Body).Decode(newUser)

		// newUser := models.User{
		// 	Name:     "apple",
		// 	Email:    "test@gmail.com",
		// 	Password: "1234",
		// }
		json.NewEncoder(rw).Encode(&newUser)
	}
}
