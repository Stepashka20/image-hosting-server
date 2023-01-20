package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (u UploadController) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	generalKey := RandKey(10)
	for _, file := range files {
		key := RandKey(10)
		filename := key + filepath.Ext(file.Filename)
		fmt.Println(filename)
		if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "upload file err"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "upload success", "key": generalKey})
}

func RandKey(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
