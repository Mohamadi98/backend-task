package service

import (
	"backend-task/database"
	"backend-task/model"
)

func CreateApplication(app model.Application) (model.Application, error) {
	result := database.Database.Create(&app)
	if result.Error != nil {
		return model.Application{}, result.Error
	}

	return app, nil
}

func GetApplications() ([]model.Application, error) {
	var app []model.Application
	result := database.Database.Find(&app)

	if result.Error != nil {
		return []model.Application{}, result.Error
	}

	return app, nil
}
