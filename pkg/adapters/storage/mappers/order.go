package mappers

import (
	"kia-logix/internal/addresses"
	"kia-logix/internal/orders"
	"kia-logix/internal/providers"
	"kia-logix/internal/user"
	"kia-logix/pkg/adapters/storage/entities"
)

func OrderDomainToEntity(domainOrder *orders.Order) *entities.Order {
	return &entities.Order{
		OwnerID:           domainOrder.OwnerID,
		SenderName:        domainOrder.Sender.Name,
		SenderPhone:       domainOrder.Sender.Phone,
		SenderAddressID:   domainOrder.Sender.AddressID,
		ReceiverName:      domainOrder.Receiver.Name,
		ReceiverPhone:     domainOrder.Receiver.Phone,
		ReceiverAddressID: domainOrder.Receiver.AddressID,
		ProviderID:        domainOrder.ProviderID,
		PickupDate:        domainOrder.PickupDate,
		Status:            string(domainOrder.Status),
	}
}

func OrderEntityToDomain(orderEntity *entities.Order) *orders.Order {
	var domainProvider *providers.Provider
	if orderEntity.Provider != nil {
		provider := ProviderEntityToDomain(*orderEntity.Provider)
		domainProvider = &provider
	}
	var domainOwner *user.User
	if orderEntity.Owner != nil {
		domainOwner = UserEntityToDomain(orderEntity.Owner)
	}
	var domainSenderAddress *addresses.Address
	if orderEntity.SenderAddress != nil {
		domainSenderAddress = AddressEntityToDomain(orderEntity.SenderAddress)
	}
	var domainReceiverAddress *addresses.Address
	if orderEntity.ReceiverAddress != nil {
		domainReceiverAddress = AddressEntityToDomain(orderEntity.ReceiverAddress)
	}
	return &orders.Order{
		ID:      orderEntity.ID,
		OwnerID: orderEntity.OwnerID,
		Owner:   domainOwner,
		Sender: &orders.Sender{
			Name:      orderEntity.SenderName,
			Phone:     orderEntity.SenderPhone,
			AddressID: orderEntity.SenderAddressID,
			Address:   domainSenderAddress,
		},
		Receiver: &orders.Receiver{
			Name:      orderEntity.ReceiverName,
			Phone:     orderEntity.ReceiverPhone,
			AddressID: orderEntity.ReceiverAddressID,
			Address:   domainReceiverAddress,
		},
		ProviderID: orderEntity.ProviderID,
		Provider:   domainProvider,
		PickupDate: orderEntity.PickupDate,
		Status:     orders.OrderStatus(orderEntity.Status),
	}
}
