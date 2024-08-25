package storage

import (
	"context"
	"gorm.io/gorm"
	"kia-logix/internal/addresses"
	"kia-logix/pkg/adapters/storage/mappers"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) addresses.Repo {
	return &addressRepo{
		db: db,
	}
}

func (o *addressRepo) Create(ctx context.Context, address *addresses.Address) (*addresses.Address, error) {
	newAddress := mappers.AddressDomainToEntity(address)
	err := o.db.WithContext(ctx).Create(&newAddress).Error
	if err != nil {
		return nil, err
	}
	createdAddress := mappers.AddressEntityToDomain(newAddress)
	return createdAddress, nil
}
