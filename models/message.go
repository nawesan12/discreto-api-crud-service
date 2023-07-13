package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID        uint
	Content       string
	ChatRoomID    uint
	PrivateChatID uint
}
