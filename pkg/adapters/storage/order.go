package storage

import (
	"context"
	"gorm.io/gorm"
	"kia-logix/internal/orders"
	"kia-logix/pkg/adapters/storage/entities"
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

func (o *orderRepo) Update(ctx context.Context, order *orders.Order) error {
	var orderEntity *entities.Order
	orderEntity = mappers.OrderDomainToEntity(order)
	if err := o.db.WithContext(ctx).Model(&orderEntity).Where("id = ?", orderEntity.ID).Updates(map[string]interface{}{"status": orderEntity.Status, "is_picked_up_sms_sent": orderEntity.IsPickedUpSMSSent, "delivered_date": orderEntity.DeliveredDate}).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) GetUpdateNeedOrders(ctx context.Context) ([]orders.Order, error) {
	var orderEntities []entities.Order
	err := o.db.WithContext(ctx).Preload("Provider").Preload("Owner").Where("status NOT IN ?", []string{string(orders.DELIVERED), string(orders.PENDING)}).Find(&orderEntities).Error
	if err != nil {
		return nil, err
	}
	orderDomainObjs := mappers.BatchOrderEntityToDomain(orderEntities)
	return orderDomainObjs, nil
}
