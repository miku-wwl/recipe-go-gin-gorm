package models

import (
	"recipe/dao"
	"time"
)

type PostCommentRequest struct {
	CommentText string `json:"comment_text"  binding:"required"`
	MediaId     int    `json:"media_id" binding:"required"`
	UserId      int    `json:"user_id"  binding:"required"`
}

type PostCommentResponse struct {
	Message string `json:"message"`
}

type Comment struct {
	CommentId   int       `json:"comment_id" gorm:"primaryKey" ` // 使用 gorm 的标签来指定主键
	MediaId     int       `json:"media_id"`
	UserId      int       `json:"user_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Comment) TableName() string {
	return "Comment"
}

func PostComment(comment Comment) error {
	err := dao.Db.Create(comment).Error
	return err
}

func GetCommentsListByMediaId(mediaId int) ([]Comment, error) {
	var commentList []Comment
	err := dao.Db.Where("media_id = ?", mediaId).Find(&commentList).Error
	return commentList, err
}
