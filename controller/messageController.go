package controller

import (
	"backend-task/model"
	"backend-task/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMessage(context *gin.Context) {
	var app model.Application
	var chat model.Chat
	var requestBody model.RequestBody
	token := context.Param("token")
	number, err := strconv.Atoi(context.Param("number"))
	if err != nil {
		fmt.Println("Conversion error: ", err)
		return
	}
	if err := context.ShouldBind(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if requestBody.Body == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Body not provided"})
		return
	}

	if err := service.GetApplicationByToken(token, &app); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.GetChatByNumber(app.ID, number, &chat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatKey := fmt.Sprintf("chat-%v-%v", chat.Number, app.Token)
	val, err := service.GetKeyInt(chatKey)
	if err != nil {
		fmt.Println("error getting chat key from redis: ", err)
		return
	}

	message := model.Message{
		Body:   requestBody.Body,
		Number: val + 1,
		ChatID: chat.ID,
	}

	if err := service.CreateMessage(&message); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.IncrementKey(chatKey); err != nil {
		fmt.Println("chat key was not inremented in redis: ", err)
	}

	if err := service.MessagesCountIncr(&chat); err != nil {
		fmt.Println("messages_count was not updated in the database: ", err)
	}

	context.JSON(http.StatusCreated, gin.H{"message": message})
}

func UpdateMessage(context *gin.Context) {
	var requestBody model.RequestBody
	var app model.Application
	var chat model.Chat
	token := context.Param("token")
	number, _ := strconv.Atoi(context.Param("number"))
	msgNumber, _ := strconv.Atoi(context.Param("msgnumber"))

	err := context.ShouldBind(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if requestBody.Body == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Body was not provided!"})
		return
	}

	if err := service.GetApplicationByToken(token, &app); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.GetChatByNumber(app.ID, number, &chat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateMessage(chat.ID, msgNumber, requestBody.Body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Message Updated Successfuly!"})
}

func GetMessages(context *gin.Context) {
	var messages []model.Message
	var app model.Application
	var chat model.Chat
	token := context.Param("token")
	number, _ := strconv.Atoi(context.Param("number"))

	if err := service.GetApplicationByToken(token, &app); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.GetChatByNumber(app.ID, number, &chat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.GetMessages(chat.ID, &messages); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Messages": messages})
}

func DeleteMessage(context *gin.Context) {
	var app model.Application
	var chat model.Chat
	token := context.Param("token")
	number, _ := strconv.Atoi(context.Param("number"))
	msgNumber, _ := strconv.Atoi(context.Param("msgnumber"))

	if err := service.GetApplicationByToken(token, &app); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.GetChatByNumber(app.ID, number, &chat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.DeleteMessage(chat.ID, msgNumber); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.MessagesCountDecr(&chat); err != nil {
		fmt.Println("messages count was not updated!")
	}

	context.JSON(http.StatusOK, gin.H{"messages": "message deleted successfuly!"})
}
