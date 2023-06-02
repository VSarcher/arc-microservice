package models

type LoginRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
