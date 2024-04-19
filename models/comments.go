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

type DeleteCommentResponse struct {
	Message string `json:"message"`
}

type UpdateCommentRequest struct {
	CommentText string `json:"comment_text"  binding:"required"`
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

func DeleteCommentById(CommentId int) error {
	var comment Comment
	err := dao.Db.Where("comment_id = ?", CommentId).Delete(&comment).Error
	return err
}

func UpdateCommentById(CommentId int, CommentText string) error {
	// 使用结构体来定义你想要更新的字段和新的值
	var data = map[string]string{
		"comment_text": CommentText, 
	}

	// 使用gorm的 Where 和 Updates 方法来更新记录
	err := dao.Db.Model(&Comment{}).Where("comment_id = ?", CommentId).Updates(data).Error
	return err
}
