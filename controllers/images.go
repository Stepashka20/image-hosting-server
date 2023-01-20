package controllers

import (
	"imagecloud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImagesController struct{}

func (i ImagesController) GetGroup(c *gin.Context) {
	key := c.Param("key")
	images, err := models.GetGroup(key)
	// delete key images.hash from each
	pathes := []string{}
	for _, image := range images {
		pathes = append(pathes, image.Path)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "group not found"})
		return
	}
	c.JSON(http.StatusOK, pathes)
}
