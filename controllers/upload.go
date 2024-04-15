package controllers

import (
	"net/http"
	"os"
	"path/filepath"
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
			Filename:  filename,
			MediaType: file.Header.Get("Content-Type"),
			Filesize:  file.Size,
		},
	}
	logger.Error(map[string]interface{}{"response": response})
	ReturnSuccess(c, 0, "success", response, 1)
}
