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

func UpdateMessage(chatID uint, number int, newBody string) error {
	var message model.Message
	err := database.Database.Where("chat_id = ? AND number = ?", chatID, number).First(&message).Error

	if err != nil {
		return err
	}

	message.Body = newBody

	if err := database.Database.Save(&message).Error; err != nil {
		return err
	}

	return nil
}

func GetMessages(chatID uint, messages *[]model.Message) error {
	err := database.Database.Where("chat_id = ?", chatID).Find(&messages).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteMessage(chatID uint, number int) error {
	err := database.Database.Where("chat_id = ? AND number = ?", chatID, number).First(&model.Message{}).
		Delete(&model.Message{}).Error

	if err != nil {
		return err
	}

	return nil
}
