package controller

import (
	"backend-task/model"
	"backend-task/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateChat(context *gin.Context) {
	var app model.Application
	token := context.Param("token")

	err := service.GetApplicationByToken(token, &app)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appKey := fmt.Sprintf("app-%v", app.Token)
	val, err := service.GetKeyInt(appKey)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	chatKey := fmt.Sprintf("chat-%v-%v", chat.Number, app.Token)

	if err := service.IncrementKey(appKey); err != nil {
		fmt.Println("app key in redis was not incremented!")
	}

	if err := service.SetKey(chatKey, 0); err != nil {
		fmt.Println("chat key was not set in redis!")
	}

	if err := service.ChatsCountIncr(&app); err != nil {
		fmt.Println("chats_count in applications table was not updated!")
	}

	context.JSON(http.StatusCreated, gin.H{"Chat": chat})
}

func GetChats(context *gin.Context) {
	var chats []model.Chat
	var app model.Application
	token := context.Param("token")

	err := service.GetApplicationByToken(token, &app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.GetChats(app.ID, &chats); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Chats": chats})
}

func GetChatByNumber(context *gin.Context) {
	var chat model.Chat
	var app model.Application
	token := context.Param("token")
	number, err := strconv.Atoi(context.Param("number"))
	if err != nil {
		fmt.Println("conversion error: ", err)
		return
	}

	err = service.GetApplicationByToken(token, &app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.GetChatByNumber(app.ID, number, &chat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Chat": chat})
}

func DeleteChat(context *gin.Context) {
	var app model.Application
	token := context.Param("token")
	number, err := strconv.Atoi(context.Param("number"))
	if err != nil {
		fmt.Println("conversion error: ", err)
		return
	}

	err = service.GetApplicationByToken(token, &app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.DeleteChat(app.ID, number); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.ChatsCountDecr(&app); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "chat deleted successfuly!"})
}
