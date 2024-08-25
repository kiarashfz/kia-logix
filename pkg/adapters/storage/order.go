package storage

import (
	"context"
	"gorm.io/gorm"
	"kia-logix/internal/orders"
	"kia-logix/pkg/adapters/storage/mappers"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) orders.Repo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(ctx context.Context, order *orders.Order) (*orders.Order, error) {
	newOrder := mappers.OrderDomainToEntity(order)
	err := o.db.WithContext(ctx).Create(&newOrder).Error
	if err != nil {
		return nil, err
	}
	createdOrder := mappers.OrderEntityToDomain(newOrder)
	return createdOrder, nil
}
