package main

import (
	"file_storage/routes"
	"file_storage/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.ConnectDB()
	utils.InitS3()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
