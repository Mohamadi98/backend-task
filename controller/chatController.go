package controller

import (
	"backend-task/model"
	"backend-task/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChat(context *gin.Context) {
	var app model.Application
	token := context.Param("token")

	err := service.GetApplicationByToken(token, &app)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if app.Token == "" {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	appKey := fmt.Sprintf("app-%v", app.Token)
	val, err := service.GetKeyInt(appKey)

	if err != nil {
		fmt.Println("error fetching key value from redis")
		return
	}

	chat := model.Chat{
		Number:        val + 1,
		ApplicationID: app.ID,
	}

	err = service.CreateChat(&chat)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.IncrementKey(appKey); err != nil {
		fmt.Println("app key in redis was not incremented!")
	}

	if err := service.ChatsCountIncr(&app); err != nil {
		fmt.Println("chats_count in applications table was not updated!")
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Chat Created Successfuly!"})
}
