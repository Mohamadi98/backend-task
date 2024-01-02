package service

import (
	"backend-task/database"
	"backend-task/model"

	"gorm.io/gorm"
)

func CreateApplication(app *model.Application) error {
	err := database.Database.Create(&app).Error
	if err != nil {
		return err
	}

	return nil
}

func GetApplications(apps *[]model.Application) error {
	err := database.Database.Find(&apps).Error

	if err != nil {
		return err
	}

	return nil
}

func GetApplicationByToken(token string, app *model.Application) error {
	err := database.Database.Where("token = ?", token).First(&app).Error

	if err != nil {
		return err
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

func DeleteApplication(token string) error {
	err := database.Database.Where("token = ?", token).First(&model.Application{}).
		Delete(&model.Application{}).Error

	if err != nil {
		return err
	}

	return nil
}

func ChatsCountIncr(app *model.Application) error {
	err := database.Database.Model(app).Update("chats_count", gorm.Expr("chats_count + ?", 1)).Error

	if err != nil {
		return err
	}

	return nil
}

func ChatsCountDecr(app *model.Application) error {
	err := database.Database.Model(app).Update("chats_count", gorm.Expr("chats_count - ?", 1)).Error

	if err != nil {
		return err
	}

	return nil
}
