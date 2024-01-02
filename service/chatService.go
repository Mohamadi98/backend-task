package service

import (
	"backend-task/database"
	"backend-task/model"
)

func CreateChat(chat *model.Chat) error {
	err := database.Database.Create(&chat).Error

	if err != nil {
		return err
	}

	return nil
}
