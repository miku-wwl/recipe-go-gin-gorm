package controllers

import (
	"net/http"
	"recipe/models"
	"recipe/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentsController struct{}

func (com CommentsController) PostComment(c *gin.Context) {
	var postCommentReq models.PostCommentRequest
	// 尝试绑定 JSON 请求体到 postCommentReq 结构体
	if err := c.ShouldBindJSON(&postCommentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Error(map[string]interface{}{"h1": postCommentReq})

	comment := models.Comment{
		MediaId:     postCommentReq.MediaId,
		UserId:      postCommentReq.UserId,
		CommentText: postCommentReq.CommentText,
		CreatedAt:   time.Now(),
	}
	logger.Info(map[string]interface{}{"Comment:": comment})

	err := models.PostComment(comment)
	if err != nil {
		ReturnError(c, 4001, "上传Comment失败, errMsg="+err.Error())
	}

	postCommentResponse := &models.PostCommentResponse{
		Message: "Comment 上传成功",
	}
	ReturnSuccess(c, 0, "success", postCommentResponse, 1)
}

func (com CommentsController) GetCommentsListByMediaId(c *gin.Context) {
	mediaIdStr := c.Param("media_id")
	mediaId, _ := strconv.Atoi(mediaIdStr)

	commentList, err := models.GetCommentsListByMediaId(mediaId)
	if err != nil {
		ReturnError(c, 4001, "获取commentList失败, errMsg="+err.Error())
		return
	}

	ReturnSuccess(c, 0, "success", commentList, 1)
}
