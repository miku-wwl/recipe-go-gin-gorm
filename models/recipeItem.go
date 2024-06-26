package models

import (
	"recipe/dao"
	"time"
)

type RecipeItem struct {
	RecipeID    int       `json:"media_id"`
	UserID      int       `json:"user_id"`
	Filename    string    `json:"filename"`
	Thumbnail   string    `json:"thumbnail"`
	FileSize    int64     `json:"filesize" gorm:"column:filesize"` // 在Go中，通常使用int64来表示文件大小
	MediaType   string    `json:"media_type"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"` // 使用指针来表示可能为nil的字符串
	Serving     string    `json:"serving"`
	CookTime    string    `json:"cook_time"`
	Ingredients string    `json:"ingredients"`
	Instruction string    `json:"instruction"`
	CreatedAt   time.Time `json:"created_at"` // 在Go中，通常使用time.Time来表示日期和时间
}

type PostRecipeRequest struct {
	// RecipeID int `json:"media_id"` --
	UserID      int     `json:"user_id"`
	Filename    string  `json:"filename"`
	Thumbnail   string  `json:"thumbnail"`
	FileSize    int64   `json:"filesize" gorm:"column:filesize"` // 在Go中，通常使用int64来表示文件大小
	MediaType   string  `json:"media_type"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"` // 使用指针来表示可能为nil的字符串
	Serving     string  `json:"serving"`
	CookTime    string  `json:"cook_time"`
	Ingredients string  `json:"ingredients"`
	Instruction string  `json:"instruction"`
	// CreatedAt time.Time `json:"created_at"` // 在Go中，通常使用time.Time来表示日期和时间 --
}

type PostRecipeResponse struct {
	Message string `json:"message"`
	Data    struct {
		RecipeID    int       `json:"media_id"`
		UserID      int       `json:"user_id"`
		Filename    string    `json:"filename"`
		Thumbnail   string    `json:"thumbnail"`
		FileSize    int64     `json:"filesize" gorm:"column:filesize"` // 在Go中，通常使用int64来表示文件大小
		MediaType   string    `json:"media_type"`
		Title       string    `json:"title"`
		Description *string   `json:"description,omitempty"` // 使用指针来表示可能为nil的字符串
		Serving     string    `json:"serving"`
		CookTime    string    `json:"cook_time"`
		Ingredients string    `json:"ingredients"`
		Instruction string    `json:"instruction"`
		CreatedAt   time.Time `json:"created_at"` // 在Go中，通常使用time.Time来表示日期和时间
	} `json:"data"`
}

func (RecipeItem) TableName() string {
	return "RecipeItem"
}

func GetRecipeItems() ([]RecipeItem, error) {
	var RecipeItems []RecipeItem
	err := dao.Db.Find(&RecipeItems).Error
	return RecipeItems, err
}

func CreateRecipeItem(recipeItem RecipeItem) (RecipeItem, error) {
	// 使用 GORM 的 Create 方法插入新数据
	if err := dao.Db.Create(&recipeItem).Error; err != nil {
		return recipeItem, err
	}
	// 返回插入的 RecipeItem
	return recipeItem, nil
}
