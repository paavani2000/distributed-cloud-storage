package routes

import (
	"file_storage/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", controllers.UploadFile)
	r.GET("/download/:id", controllers.DownloadFile)
}
