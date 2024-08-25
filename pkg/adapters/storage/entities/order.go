package entities

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OwnerID           uint
	Owner             *User  `gorm:"foreignKey:OwnerID"`
	SenderName        string `gorm:"not null; type:varchar(255)"`
	SenderPhone       string `gorm:"not null; type:varchar(15)"`
	SenderAddressID   uint
	SenderAddress     *Address `gorm:"foreignKey:SenderAddressID"`
	ReceiverName      string   `gorm:"not null; type:varchar(255)"`
	ReceiverPhone     string   `gorm:"not null; type:varchar(15)"`
	ReceiverAddressID uint
	ReceiverAddress   *Address `gorm:"foreignKey:ReceiverAddressID"`
	ProviderID        uint
	Provider          *Provider `gorm:"foreignKey:ProviderID"`
	PickupDate        *time.Time
	DeliveredDate     *time.Time
	Status            string `gorm:"not null; type:varchar(63)"`
	IsPickedUpSMSSent bool   `gorm:"default:false"`
}
