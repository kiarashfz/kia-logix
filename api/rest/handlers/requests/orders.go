package requests

import (
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/internal/orders"
)

type CreateOrder struct {
	Sender     *helpers.Sender   `json:"sender" validate:"required"`
	Receiver   *helpers.Receiver `json:"Receiver" validate:"required"`
	ProviderID uint              `json:"provider_id" validate:"required"`
	PickupDate helpers.Date      `json:"pickup_date" validate:"required"`
}

func CreateOrderToDomainOrder(o *CreateOrder) *orders.Order {
	return &orders.Order{
		Sender: &orders.Sender{
			Name:      o.Sender.Name,
			Phone:     o.Sender.Phone,
			AddressID: o.Sender.AddressID,
		},
		Receiver: &orders.Receiver{
			Name:      o.Receiver.Name,
			Phone:     o.Receiver.Phone,
			AddressID: o.Receiver.AddressID,
		},
		ProviderID: o.ProviderID,
		PickupDate: o.PickupDate.Time,
	}
}
