package main

import (
	"log"
	"snippet-saver/internal"
	"snippet-saver/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	database.InitDb()
	defer database.Close()

	router := gin.Default()
	// router.Use(middleware.CORSMiddleware())
	internal.InitializeRoutes(router)
	err := router.Run(":8080")

	if err != nil {
		log.Fatal("Failed to run server", err)
	}

}
