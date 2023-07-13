package models

import (
	"gorm.io/gorm"
)

type ChatRoom struct {
	gorm.Model
	Name       string
	CompanyID  uint
	Password   string
	Users      []User `gorm:"many2many:user_chatrooms;"`
	Messages   []Message
}
