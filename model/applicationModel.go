package model

type Application struct {
	ID         uint   `gorm:"primaryKey" json:"-"`
	Token      string `gorm:"size:255;not null;unique" json:"token"`
	Name       string `gorm:"size:255;not null;unique" json:"name"`
	ChatsCount int    `gorm:"column:chats_count" json:"chats_count"`
	Chat       []Chat `gorm:"constraint:OnDelete:CASCADE;"`
}
