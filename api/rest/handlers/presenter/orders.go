package presenter

import (
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/internal/orders"
)

type OrderResp struct {
	ID         uint              `json:"id"`
	Sender     *helpers.Sender   `json:"sender"`
	Receiver   *helpers.Receiver `json:"receiver"`
	ProviderID uint              `json:"provider_id"`
	PickupDate helpers.Date      `json:"pickup_date"`
	Status     string            `json:"status"`
}

func OrderToCreateOrderResp(domainOrder *orders.Order) *OrderResp {
	return &OrderResp{
		ID:         domainOrder.ID,
		Sender:     helpers.SenderToSenderResp(domainOrder.Sender),
		Receiver:   helpers.ReceiverToReceiverResp(domainOrder.Receiver),
		ProviderID: domainOrder.ProviderID,
		PickupDate: helpers.Date{
			Time: domainOrder.PickupDate,
		},
		Status: string(domainOrder.Status),
	}
}
