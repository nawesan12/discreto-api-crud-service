package models

import (
	"gorm.io/gorm"
)

type PrivateChat struct {
	gorm.Model
	UserOneID uint
	UserTwoID uint
	Messages  []Message
}
