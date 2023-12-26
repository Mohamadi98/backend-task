package controller

import (
	"backend-task/model"
	"backend-task/service"
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

	_, err := service.CreateApplication(app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Application Created Successfuly!"})
}

func GetApplications(context *gin.Context) {
	apps, err := service.GetApplications()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"applications": apps})
}

func GetApplicationByToken(context *gin.Context) {
	token := context.Param("token")

	app, err := service.GetApplicationByToken(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if app.Token == "" {
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

	IsUpdated, err := service.UpdateApplication(token, requestData.Name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !IsUpdated {
		context.JSON(http.StatusNotFound, gin.H{"error": "No Row Found With This Token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Application Updated Successfuly!"})
}

func DeleteApplication(context *gin.Context) {
	token := context.Param("token")

	IsUpdated, err := service.DeleteApplication(token)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !IsUpdated {
		context.JSON(http.StatusNotFound, gin.H{"error": "No Row Found With This Token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Application Deleted Successfuly!"})
}
