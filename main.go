package main

import (
	"backend-task/controller"
	"backend-task/database"
	"backend-task/model"
	"backend-task/redis"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	loadRedis()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.GET("/", serverHealth)
	appRoutes := router.Group("/application")
	appRoutes.POST("", controller.CreateApplication)
	appRoutes.GET("", controller.GetApplications)
	appRoutes.GET("/:token", controller.GetApplicationByToken)
	appRoutes.PUT("/:token", controller.UpdateApplication)
	appRoutes.DELETE("/:token", controller.DeleteApplication)
	appRoutes.POST("/:token/chat", controller.CreateChat)
	appRoutes.GET("/:token/chat", controller.GetChats)
	appRoutes.GET("/:token/chat/:number", controller.GetChatByNumber)
	appRoutes.DELETE("/:token/chat/:number", controller.DeleteChat)
	appRoutes.POST("/:token/chat/:number", controller.CreateMessage)
	appRoutes.PUT("/:token/chat/:number/message/:msgnumber", controller.UpdateMessage)
	appRoutes.GET("/:token/chat/:number/message", controller.GetMessages)
	appRoutes.DELETE("/:token/chat/:number/message/:msgnumber", controller.DeleteMessage)

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

func loadRedis() {
	redis.Connect()
}

func serverHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server Running!"})
}
