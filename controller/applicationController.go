package controller

import (
	"backend-task/model"
	"backend-task/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateApplication(context *gin.Context) {
	var requestData model.RequestBody

	if err := context.ShouldBind(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if requestData.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Name not provided"})
		return
	}

	app := model.Application{
		Name:  requestData.Name,
		Token: service.GenerateIdentifier(),
	}

	err := service.CreateApplication(&app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appKey := fmt.Sprintf("app-%v", app.Token)
	redisError := service.SetKey(appKey, 0)
	if redisError != nil {
		fmt.Println("error creating an app key: ", redisError)
	}

	context.JSON(http.StatusCreated, gin.H{"Application": app})
}

func GetApplications(context *gin.Context) {
	var apps []model.Application
	err := service.GetApplications(&apps)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"applications": apps})
}

func GetApplicationByToken(context *gin.Context) {
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

	context.JSON(http.StatusOK, gin.H{"application": app})
}

func UpdateApplication(context *gin.Context) {
	token := context.Param("token")
	var requestData model.RequestBody

	if err := context.ShouldBind(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if requestData.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Name Not Provided"})
		return
	}

	isUpdated, err := service.UpdateApplication(token, requestData.Name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !isUpdated {
		context.JSON(http.StatusNotFound, gin.H{"error": "No Row Found With This Token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Application Updated Successfuly!"})
}

func DeleteApplication(context *gin.Context) {
	token := context.Param("token")

	isDeleted, err := service.DeleteApplication(token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !isDeleted {
		context.JSON(http.StatusNotFound, gin.H{"error": "No Row Found With This Token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Application Deleted Successfuly!"})
}
