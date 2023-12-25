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

	createdApp, err := service.CreateApplication(app)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"application": createdApp})
}

func GetApplications(context *gin.Context) {
	apps, err := service.GetApplications()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"applications": apps})
}
