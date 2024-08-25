package service

import (
	"context"
	"kia-logix/factory"
	"kia-logix/internal/orders"
	"kia-logix/pkg/adapters/sms"
	"time"
)

type IOrderService interface {
	Create(ctx context.Context, order *orders.Order, userID uint) (*orders.Order, error)
	GetUpdateNeedOrders(ctx context.Context) ([]orders.Order, error)
	UpdateOrderStatusJob(ctx context.Context, order *orders.Order) error

	// Tasks
	UpdateOrderStatusTask(ctx context.Context, args ...interface{}) error
}

type OrderService struct {
	orderOps   orders.IOrderOps
	SMSService sms.ISMSService
}

func NewOrderService(orderOps orders.IOrderOps, smsService sms.ISMSService) *OrderService {
	return &OrderService{orderOps: orderOps, SMSService: smsService}
}

func (o *OrderService) Create(ctx context.Context, order *orders.Order, userID uint) (*orders.Order, error) {
	order.Status = orders.REGISTERED
	order.OwnerID = userID
	return o.orderOps.Create(ctx, order)
}

func (o *OrderService) UpdateOrderStatusJob(ctx context.Context, order *orders.Order) error {
	statusProvider, err := factory.GetStatusProvider(order.Provider.Name, order.Provider.URL)
	if err != nil {
		return err
	}
	newStatus, err := statusProvider.GetOrderStatus(order.ID)
	if err != nil {
		return err
	}
	order.Status = newStatus
	if newStatus == orders.PICKED_UP && !order.IsPickedUpSMSSent {
		go o.SMSService.Send(order.Owner.Phone, "custom template")
		order.IsPickedUpSMSSent = true
	}
	if newStatus == orders.DELIVERED && order.DeliveredDate == nil {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		order.DeliveredDate = &today
	}
	return o.orderOps.Update(ctx, order)
}

func (o *OrderService) GetUpdateNeedOrders(ctx context.Context) ([]orders.Order, error) {
	// order status not in PENDING, DELIVERED
	return o.orderOps.GetUpdateNeedOrders(ctx)
}
