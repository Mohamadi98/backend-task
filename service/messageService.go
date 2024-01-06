package service

import (
	"backend-task/database"
	"backend-task/model"
)

func CreateMessage(message *model.Message) error {
	err := database.Database.Create(&message).Error

	if err != nil {
		return err
	}

	return nil
}
