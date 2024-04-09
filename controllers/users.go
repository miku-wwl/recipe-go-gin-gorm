package controllers

import (
	"recipe/models"
	"strconv"

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
