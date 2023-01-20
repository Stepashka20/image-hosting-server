package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"

	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"imagecloud/models"
)

type UploadController struct{}

func (u UploadController) Upload(c *gin.Context) {
	fmt.Println("upload")
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	generalKey := RandKey(10)
	allKeys := []string{}
	for _, file := range files {
		hash := CalcHash(file)
		fmt.Println(hash)
		image := models.Image{}.GetImageByHash(hash)
		fmt.Println(image)
		if image.Key != "" {
			allKeys = append(allKeys, image.Key)
			continue
		}

		key := RandKey(10)
		filename := key + filepath.Ext(file.Filename)

		if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "upload file err"})
			return
		}
		image = models.Image{}.NewImage(key, hash, filename)
		allKeys = append(allKeys, image.Key)
	}
	fmt.Println(allKeys)
	models.Image{}.NewGroupImages(generalKey, allKeys)
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

func CalcHash(file *multipart.FileHeader) string {
	hash := sha256.New()
	f, _ := file.Open()
	io.Copy(hash, f)
	return hex.EncodeToString(hash.Sum(nil))
}
