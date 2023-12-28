package service

import (
	"backend-task/database"
	"backend-task/model"
)

func CreateApplication(app *model.Application) error {
	result := database.Database.Create(&app)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetApplications(apps *[]model.Application) error {
	result := database.Database.Find(&apps)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetApplicationByToken(token string, app *model.Application) error {
	result := database.Database.Where("token = ?", token).Find(&app)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateApplication(token, newName string) (bool, error) {
	result := database.Database.Model(&model.Application{}).Where("token = ?", token).
		Update("name", newName)

	if result.Error != nil || result.RowsAffected == 0 {
		return false, result.Error
	}

	return true, nil
}

func DeleteApplication(token string) (bool, error) {
	result := database.Database.Delete(&model.Application{}, "token = ?", token)

	if result.Error != nil || result.RowsAffected == 0 {
		return false, result.Error
	}

	return true, nil
}
