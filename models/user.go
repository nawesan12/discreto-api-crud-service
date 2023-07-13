package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string    `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(100)" json:"password"`
	FirstName string    `gorm:"type:varchar(100)" json:"first_name,omitempty"`
	LastName  string    `gorm:"type:varchar(100)" json:"last_name,omitempty"`
	CompanyID string    `gorm:"type:varchar(100)" json:"company_id,omitempty"`
	IsAdmin   bool      `gorm:"type:bool" json:"is_admin,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
