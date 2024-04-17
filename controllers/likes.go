package controllers

import (
	"net/http"
	"recipe/models"
	"recipe/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LikesController struct{}

func (l LikesController) PostLike(c *gin.Context) {
	var postLikeReq models.PostLikeRequest
	// 尝试绑定 JSON 请求体到 postLikeReq 结构体
	if err := c.ShouldBindJSON(&postLikeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Error(map[string]interface{}{"h1": postLikeReq})

	like := models.Like{
		MediaId:   postLikeReq.MediaId,
		UserId:    postLikeReq.UserId,
		CreatedAt: time.Now(),
	}
	logger.Info(map[string]interface{}{"like:": like})

	err := models.PostLike(like)
	if err != nil {
		ReturnError(c, 4001, "上传Like失败, errMsg="+err.Error())
	}

	postLikeResponse := &models.PostLikeResponse{
		Message: "Like 上传成功",
	}
	ReturnSuccess(c, 0, "success", postLikeResponse, 1)
}

func (l LikesController) DeleteLike(c *gin.Context) {
	likeIdStr := c.Param("like_id")
	likeId, _ := strconv.Atoi(likeIdStr)
	err := models.DeleteLike(likeId)

	if err != nil {
		ReturnError(c, 4001, "删除Like失败, errMsg="+err.Error())
		return
	}

	deleteLikeResponse := &models.DeleteLikeResponse{
		Message: "Like 删除成功",
	}
	ReturnSuccess(c, 0, "success", deleteLikeResponse, 1)
}

func (l LikesController) GetCountByMediaId(c *gin.Context) {
	mediaIdStr := c.Param("media_id")
	mediaId, _ := strconv.Atoi(mediaIdStr)
	count, err := models.GetCountByMediaId(mediaId)
	if err != nil {
		ReturnError(c, 4001, "获取Like统计失败, errMsg="+err.Error())
		return
	}

	deleteLikeResponse := &models.GetCountResponse{
		Count: count,
	}
	ReturnSuccess(c, 0, "success", deleteLikeResponse, 1)
}

func (l LikesController) GetUserLike(c *gin.Context) {
	var getUserLikeReq models.GetUserLikeRequest
	// 尝试绑定 JSON 请求体到 getUserLikeReq 结构体
	if err := c.ShouldBindJSON(&getUserLikeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	like := models.Like{
		MediaId: getUserLikeReq.MediaId,
		UserId:  getUserLikeReq.UserId,
	}
	logger.Info(map[string]interface{}{"like:": like})

	like, err := models.GetUserLike(getUserLikeReq.MediaId, getUserLikeReq.UserId)
	if err != nil {
		ReturnError(c, 4001, "获取用户Like失败, errMsg="+err.Error())
	}

	ReturnSuccess(c, 0, "success", like, 1)
}
