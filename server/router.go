package server

import (
	"imagecloud/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	// "imagecloud/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(static.Serve("/img", static.LocalFile("./uploads", true)))
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	api := router.Group("/api")
	{
		// upload
		upload := new(controllers.UploadController)
		api.POST("/upload", upload.Upload)
	}

	return router

}
