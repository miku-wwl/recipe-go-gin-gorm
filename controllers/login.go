package controllers

import (
	"recipe/models"
	"recipe/pkg/jwtServer"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (l LoginController) GetLoginResponse(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
	}

	user, _ := models.GetUserInfoByUserName(username)
	if user.UserID == 0 {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}
	if user.Password != password {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}

	token, err := jwtServer.CreateToken(int64(user.UserID))
	c.Set("token", token)
	levelNames := map[int]string{
		1: "admin",
		2: "user",
		3: "guest",
	}

	data := &models.LoginResponse{
		Token:   token,
		Message: "登录成功并返回token成功",
		User: models.UserWithNoPassword{
			UserID:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Levelname: levelNames[user.UserLevelID],
		},
	}

	if err == nil {
		ReturnSuccess(c, 0, "success", data, 1)
		return
	}
	ReturnError(c, 4001, "创建token失败")
}
