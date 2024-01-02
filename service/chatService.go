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

func GetChats(appID uint, chats *[]model.Chat) error {
	err := database.Database.Where("app_id = ?", appID).Find(&chats).Error

	if err != nil {
		return err
	}

	return nil
}

func GetChatByNumber(appID uint, number int, chat *model.Chat) error {
	err := database.Database.Where("app_id = ? AND number = ?", appID, number).First(&chat).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteChat(appID uint, number int) error {
	err := database.Database.Where("app_id = ? AND number = ?", appID, number).First(&model.Chat{}).
		Delete(&model.Chat{}).Error

	if err != nil {
		return err
	}

	return nil
}
