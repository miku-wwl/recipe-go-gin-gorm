package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"recipe/models"
	"recipe/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadResponse struct {
	Message string `json:"message"`
	Data    struct {
		Filename  string `json:"filename"`
		MediaType string `json:"media_type"`
		Filesize  int64  `json:"filesize"`
	} `json:"data"`
}

type UploadController struct{}

func (p UploadController) UploadFiles(c *gin.Context) {
	logger.Error(map[string]interface{}{"hello": "hello"})
	file, err := c.FormFile("file")
	logger.Error(map[string]interface{}{"file": file})
	if err != nil {
		c.JSON(http.StatusBadRequest, UploadResponse{Message: "No file uploaded"})
		return
	}
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + file.Filename
	filenameRes := filename
	// 设置保存路径
	uploadDir := "/var/www/html/recipe/images"
	filename = filepath.Join(uploadDir, filename)
	logger.Error(map[string]interface{}{"filename": filename})
	// 创建目录（如果不存在）
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// 保存文件
	logger.Error(map[string]interface{}{"err": err})
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Error(map[string]interface{}{"err": err})

	logger.Error(map[string]interface{}{"err": err})
	response := UploadResponse{
		Message: "File uploaded successfully",
		Data: struct {
			Filename  string `json:"filename"`
			MediaType string `json:"media_type"`
			Filesize  int64  `json:"filesize"`
		}{
			Filename:  filenameRes,
			MediaType: file.Header.Get("Content-Type"),
			Filesize:  file.Size,
		},
	}
	logger.Error(map[string]interface{}{"response": response})
	ReturnSuccess(c, 0, "success", response, 1)
}

func NewRecipeItemFromRequest(postReq models.PostRecipeRequest) models.RecipeItem {
	return models.RecipeItem{
		UserID:      postReq.UserID,
		Filename:    postReq.Filename,
		Thumbnail:   postReq.Thumbnail,
		FileSize:    postReq.FileSize,
		MediaType:   postReq.MediaType,
		Title:       postReq.Title,
		Description: postReq.Description,
		Serving:     postReq.Serving,
		CookTime:    postReq.CookTime,
		Ingredients: postReq.Ingredients,
		Instruction: postReq.Instruction,
		CreatedAt:   time.Now(), // 设置创建时间为当前时间
	}
}

func BuildPostRecipeResponse(retRecipeItem models.RecipeItem) models.PostRecipeResponse {
	return models.PostRecipeResponse{
		Message: "Recipe uploaded successfully",
		Data: struct {
			RecipeID    int       `json:"media_id"`
			UserID      int       `json:"user_id"`
			Filename    string    `json:"filename"`
			Thumbnail   string    `json:"thumbnail"`
			FileSize    int64     `json:"filesize" gorm:"column:filesize"`
			MediaType   string    `json:"media_type"`
			Title       string    `json:"title"`
			Description *string   `json:"description,omitempty"`
			Serving     string    `json:"serving"`
			CookTime    string    `json:"cook_time"`
			Ingredients string    `json:"ingredients"`
			Instruction string    `json:"instruction"`
			CreatedAt   time.Time `json:"created_at"`
		}{
			RecipeID:    retRecipeItem.RecipeID,
			UserID:      retRecipeItem.UserID,
			Filename:    retRecipeItem.Filename,
			Thumbnail:   retRecipeItem.Thumbnail,
			FileSize:    retRecipeItem.FileSize,
			MediaType:   retRecipeItem.MediaType,
			Title:       retRecipeItem.Title,
			Description: retRecipeItem.Description,
			Serving:     retRecipeItem.Serving,
			CookTime:    retRecipeItem.CookTime,
			Ingredients: retRecipeItem.Ingredients,
			Instruction: retRecipeItem.Instruction,
			CreatedAt:   retRecipeItem.CreatedAt,
		},
	}
}

func (p UploadController) PostRecipe(c *gin.Context) {
	var postReq models.PostRecipeRequest
	// 尝试绑定 JSON 请求体到 loginReq 结构体
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipeItem := NewRecipeItemFromRequest(postReq)

	retRecipeItem, err := models.CreateRecipeItem(recipeItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := BuildPostRecipeResponse(retRecipeItem)

	ReturnSuccess(c, 0, "success", response, 1)
}
