package models

import (
	"recipe/dao"
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
	UserID      int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	UserLevelID int       `json:"user_level_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserWithNoPassword struct {
	UserID    int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Levelname string    `json:"level_name"`
}

func (User) TableName() string {
	return "User"
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
