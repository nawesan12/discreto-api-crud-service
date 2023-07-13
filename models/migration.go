package models

import (
	"gorm.io/gorm"
)

func MigrateTables(DB *gorm.DB) {
	DB.AutoMigrate(&User{}, &Company{}, &ChatRoom{}, &PrivateChat{}, &Message{})
}
