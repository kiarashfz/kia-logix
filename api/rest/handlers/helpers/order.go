package helpers

import "kia-logix/internal/orders"

type Sender struct {
	Name      string `json:"name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	AddressID uint   `json:"address_id" validate:"required"`
}

type Receiver struct {
	Name      string `json:"name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	AddressID uint   `json:"address_id" validate:"required"`
}

func SenderToSenderResp(domainSender *orders.Sender) *Sender {
	return &Sender{
		Name:      domainSender.Name,
		Phone:     domainSender.Phone,
		AddressID: domainSender.AddressID,
	}
}

func ReceiverToReceiverResp(domainReceiver *orders.Receiver) *Receiver {
	return &Receiver{
		Name:      domainReceiver.Name,
		Phone:     domainReceiver.Phone,
		AddressID: domainReceiver.AddressID,
	}
}
