package model

type Chat struct {
	ID            uint      `gorm:"primaryKey" json:"-"`
	Number        uint      `json:"number"`
	MessagesCount uint      `gorm:"column:messages_count" json:"messages_count"`
	Message       []Message `gorm:"constraint:OnDelete:CASCADE;"`
	ApplicationID uint      `gorm:"column:app_id" json:"app_id"`
}
