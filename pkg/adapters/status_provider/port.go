package status_provider

import "kia-logix/internal/orders"

type IStatusProvider interface {
	GetOrderStatus(orderID uint) (orders.OrderStatus, error)
}
