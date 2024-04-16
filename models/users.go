package models

import (
	"recipe/dao"
	"recipe/pkg/logger"
	"time"
)

type User struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	Email       string    `json:"email"`
	UserLevelID int       `json:"user_level_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserApi struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	UserLevelID int       `json:"user_level_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserWithNoPassword struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Levelname string    `json:"level_name"`
}

type UserResponse struct {
	Message string             `json:"message"`
	User    UserWithNoPassword `json:"user"`
}

type PostUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AvailableResponse struct {
	Available bool `json:"available"`
}

func (User) TableName() string {
	return "User"
}

func GetUsernameAvailable(username string) (AvailableResponse, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error(map[string]interface{}{"GetUsernameAvailable error": err})
	}
	if user.UserID == 0 {
		return AvailableResponse{Available: true}, err
	}
	return AvailableResponse{Available: false}, err
}

func GetEmailAvailable(email string) (AvailableResponse, error) {
	var user User
	err := dao.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Error(map[string]interface{}{"GetEmailAvailable error": err})
	}
	if user.UserID == 0 {
		return AvailableResponse{Available: true}, err
	}
	return AvailableResponse{Available: false}, err
}

func GetUserByUserId(id int) (UserApi, error) {
	var user User
	err := dao.Db.Where("user_id = ?", id).First(&user).Error
	userApi := UserApi{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		UserLevelID: user.UserLevelID,
		CreatedAt:   user.CreatedAt,
	}
	return userApi, err
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func PostUser(user User) error {
	err := dao.Db.Create(user).Error
	return err
}
