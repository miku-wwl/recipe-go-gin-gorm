package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wangwu")

	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)

	search := &Search{}
	err := c.Bind(&search)

	if err == nil {
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
}
