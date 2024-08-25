package entities

import (
	"gorm.io/gorm"
	"kia-logix/internal/addresses"
)

type Address struct {
	gorm.Model
	State       string                 `gorm:"type:varchar(31);not null"`
	City        string                 `gorm:"type:varchar(31);not null"`
	PostalCode  string                 `gorm:"type:varchar(31)"`
	AddressLine string                 `gorm:"not null"`
	Coordinates *addresses.Coordinates `gorm:"type:geography(POINT, 4326)"`
	Owner       *User                  `gorm:"foreignKey:OwnerID;not null"`
	OwnerID     uint                   `gorm:"not null"`
}
