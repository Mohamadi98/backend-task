package main

import (
	"backend-task/controller"
	"backend-task/database"
	"backend-task/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.GET("/", serverHealth)
	appRoutes := router.Group("/application")
	appRoutes.POST("", controller.CreateApplication)
	appRoutes.GET("", controller.GetApplications)

	router.Run("localhost:8080")
	fmt.Println("Server Running On Port 8080")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Application{}, &model.Chat{}, &model.Message{})
}

func serverHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server Running!"})
}
