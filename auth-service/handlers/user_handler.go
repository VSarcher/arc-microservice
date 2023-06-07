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
	ERR_DECODING_JSON    = "Error on decoding json request"
	ERR_RECORD_NOT_FOUND = "Error on retrieving data from table, Not Found"
)

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Create New User
		newUser := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			JSON_response(rw, http.StatusInternalServerError, map[string]string{"error": ERR_DECODING_JSON})
			log.Println(ERR_DECODING_JSON)
			return
		}

		// fmt.Println(newUser)
		config.PostgresDB.DB.Create(&newUser)

		// rw.Header().Add("Content-Type", "application/json")
		// json.NewEncoder(rw).Encode(newUser)
		JSON_response(rw, http.StatusOK, newUser)
	}
}

func LoginUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userInfo := new(models.LoginRequestModel)
		if err := json.NewDecoder(r.Body).Decode(userInfo); err != nil {
			JSON_response(rw, http.StatusInternalServerError, map[string]string{"error": ERR_DECODING_JSON})
			log.Println(ERR_DECODING_JSON)
			return
		}

		// log.Printf("----------------%s\n", userInfo)
		user := new(models.User)
		err := config.PostgresDB.DB.Where("email = ?", userInfo.Email).First(&user)
		log.Println(err)
		if err != nil {
			JSON_response(rw, http.StatusInternalServerError, map[string]string{"error": ERR_RECORD_NOT_FOUND})
			log.Println(ERR_RECORD_NOT_FOUND)
			return
		}
		JSON_response(rw, http.StatusOK, user)
	}
}
