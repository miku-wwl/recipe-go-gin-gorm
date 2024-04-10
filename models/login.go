package models

type LoginResponse struct {
	Token   string             `json:"token"`
	Message string             `json:"message"`
	User    UserWithNoPassword `json:"user"`
}
