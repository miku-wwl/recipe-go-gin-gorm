package controllers

import (
	"net/http"
	"recipe/models"
	"recipe/pkg/jwtServer"
	"recipe/pkg/logger"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (l LoginController) GetLoginResponse(c *gin.Context) {
	var loginReq models.LoginRequest
	// 尝试绑定 JSON 请求体到 loginReq 结构体
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := loginReq.Username
	password := loginReq.Password
	logger.Info(map[string]interface{}{"username:": username})
	logger.Info(map[string]interface{}{"password:": password})

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
	retToken := ""
	session := sessions.Default(c)
	if token, ok := session.Get(strconv.Itoa(user.UserID) + " :token").(string); ok {
		retToken = token
		// c.JSON(200, gin.H{"token": token})
	} else {
		// token不存在
		newToken, err := jwtServer.CreateToken(int64(user.UserID))
		retToken = newToken
		if err != nil {
			ReturnError(c, 4001, "创建token失败")
			return
		}
		session.Set(strconv.Itoa(user.UserID)+" :token", retToken)
		session.Save()
	}

	levelNames := map[int]string{
		1: "admin",
		2: "user",
		3: "guest",
	}

	data := &models.LoginResponse{
		Token:   retToken,
		Message: "登录成功并返回token成功",
		User: models.UserWithNoPassword{
			UserID:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Levelname: levelNames[user.UserLevelID],
		},
	}
	ReturnSuccess(c, 0, "success", data, 1)
}
