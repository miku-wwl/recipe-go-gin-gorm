package controllers

import (
	"net/http"
	"recipe/models"
	"recipe/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (u UsersController) GetUserByUserId(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	userApi, err := models.GetUserByUserId(id)

	if err != nil {
		ReturnError(c, 4001, "获取失败")
		return
	}
	ReturnSuccess(c, 0, "success", userApi, 1)
}

func (u UsersController) GetUsernameAvailable(c *gin.Context) {
	username := c.Param("username")
	available, err := models.GetUsernameAvailable(username)
	logger.Error(map[string]interface{}{"UsersController GetUsernameAvailable error": err})

	if err != nil && err.Error() != "record not found" {
		ReturnError(c, 4001, "获取username失败")
		return
	}
	ReturnSuccess(c, 0, "success", available, 1)
}

func (u UsersController) GetEmailAvailable(c *gin.Context) {
	email := c.Param("email")
	available, err := models.GetEmailAvailable(email)
	logger.Error(map[string]interface{}{"UsersController GetEmailAvailable error": err})

	if err != nil && err.Error() != "record not found" {
		ReturnError(c, 4001, "获取username失败")
		return
	}
	ReturnSuccess(c, 0, "success", available, 1)
}

func (u UsersController) PostUser(c *gin.Context) {
	var postUserReq models.PostUserRequest
	// 尝试绑定 JSON 请求体到 postUserReq 结构体
	if err := c.ShouldBindJSON(&postUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := postUserReq.Username
	password := postUserReq.Password
	email := postUserReq.Email
	logger.Info(map[string]interface{}{"username:": username})
	logger.Info(map[string]interface{}{"password:": password})
	logger.Info(map[string]interface{}{"email:": email})
	user := models.User{
		Username:    username,
		Password:    password,
		Email:       email,
		UserLevelID: 2, // 设置用户级别为2 user
		CreatedAt:   time.Now(),
	}

	err := models.PostUser(user)
	if err != nil {
		ReturnError(c, 4001, "注册失败，errMsg="+err.Error())
	}

	userTemp, err := models.GetUserInfoByUserName(user.Username)
	if err != nil {
		logger.Error(map[string]interface{}{"[PostUser] GetUserInfoByUserName error": err})
	}
	data := &models.UserResponse{
		Message: "注册成功",
		User: models.UserWithNoPassword{
			UserID:    userTemp.UserID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Levelname: "user",
		},
	}
	ReturnSuccess(c, 0, "success", data, 1)
}
