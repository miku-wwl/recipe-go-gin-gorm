package controllers

import (
	"recipe/models"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

func (p RecipeController) GetRecipeItems(c *gin.Context) {
	rs, err := models.GetRecipeItems()
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}
