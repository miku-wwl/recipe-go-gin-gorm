package models

type LoginResponse struct {
	Token   string             `json:"token"`
	Message string             `json:"message"`
	User    UserWithNoPassword `json:"user"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
