package model

type Message struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Number uint   `json:"number"`
	Body   string `json:"body"`
	ChatID uint   `gorm:"column:chat_id" json:"chat_id"`
}
