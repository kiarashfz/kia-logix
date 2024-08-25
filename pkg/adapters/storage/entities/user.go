package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"unique;type:varchar(15);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `gorm:"not null"`
}
