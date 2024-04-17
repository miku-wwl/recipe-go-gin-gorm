package models

import (
	"recipe/dao"
	"time"
)

type Like struct {
	LikeID    int       `json:"like_id" gorm:"primaryKey" ` // 使用 gorm 的标签来指定主键
	MediaId   int       `json:"media_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (Like) TableName() string {
	return "Like"
}

type PostLikeRequest struct {
	MediaId int `json:"media_id" binding:"required"`
	UserId  int `json:"user_id"  binding:"required"`
}

type GetUserLikeRequest struct {
	MediaId int `json:"media_id" binding:"required"`
	UserId  int `json:"user_id"  binding:"required"`
}

type PostLikeResponse struct {
	Message string `json:"message"`
}

type DeleteLikeResponse struct {
	Message string `json:"message"`
}

type GetCountResponse struct {
	Count int `json:"count"`
}

func PostLike(like Like) error {
	err := dao.Db.Create(like).Error
	return err
}

func DeleteLike(likeId int) error {
	var like Like
	err := dao.Db.Where("like_id = ?", likeId).Delete(&like).Error
	return err
}

func GetCountByMediaId(mediaId int) (int, error) {
	var count int
	err := dao.Db.Model(&Like{}).Where("media_id = ?", mediaId).Count(&count).Error
	return count, err
}

func GetUserLike(mediaId int, userId int) (Like, error) {
	var like Like
	err := dao.Db.Where("media_id = ?", mediaId).Where("user_id = ?", userId).First(&like).Error
	return like, err
}
