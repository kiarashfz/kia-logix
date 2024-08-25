package service

import (
	"context"
	"kia-logix/internal/orders"
)

type IOrderService interface {
	Create(ctx context.Context, order *orders.Order, userID uint) (*orders.Order, error)
}

type OrderService struct {
	orderOps orders.IOrderOps
}

func NewOrderService(orderOps orders.IOrderOps) *OrderService {
	return &OrderService{orderOps: orderOps}
}

func (o *OrderService) Create(ctx context.Context, order *orders.Order, userID uint) (*orders.Order, error) {
	order.Status = orders.REGISTERED
	order.OwnerID = userID
	return o.orderOps.Create(ctx, order)
}
