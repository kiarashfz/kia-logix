package mappers

import (
	"gorm.io/gorm"
	"kia-logix/internal/addresses"
	"kia-logix/internal/orders"
	"kia-logix/internal/providers"
	"kia-logix/internal/user"
	"kia-logix/pkg/adapters/storage/entities"
)

func OrderDomainToEntity(domainOrder *orders.Order) *entities.Order {
	var (
		senderName        string
		senderPhone       string
		senderAddressID   uint
		receiverName      string
		receiverPhone     string
		receiverAddressID uint
	)
	if domainOrder.Sender != nil {
		senderName = domainOrder.Sender.Name
		senderPhone = domainOrder.Sender.Phone
		senderAddressID = domainOrder.Sender.AddressID
	}
	if domainOrder.Receiver != nil {
		receiverName = domainOrder.Receiver.Name
		receiverPhone = domainOrder.Receiver.Phone
		receiverAddressID = domainOrder.Receiver.AddressID
	}
	return &entities.Order{
		Model:             gorm.Model{ID: domainOrder.ID},
		OwnerID:           domainOrder.OwnerID,
		SenderName:        senderName,
		SenderPhone:       senderPhone,
		SenderAddressID:   senderAddressID,
		ReceiverName:      receiverName,
		ReceiverPhone:     receiverPhone,
		ReceiverAddressID: receiverAddressID,
		ProviderID:        domainOrder.ProviderID,
		PickupDate:        domainOrder.PickupDate,
		DeliveredDate:     domainOrder.DeliveredDate,
		Status:            string(domainOrder.Status),
		IsPickedUpSMSSent: domainOrder.IsPickedUpSMSSent,
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
		ProviderID:        orderEntity.ProviderID,
		Provider:          domainProvider,
		PickupDate:        orderEntity.PickupDate,
		Status:            orders.OrderStatus(orderEntity.Status),
		IsPickedUpSMSSent: orderEntity.IsPickedUpSMSSent,
	}
}

func BatchOrderEntityToDomain(orderEntities []entities.Order) []orders.Order {
	var domainOrders []orders.Order
	for _, orderEntity := range orderEntities {
		domainOrders = append(domainOrders, *OrderEntityToDomain(&orderEntity))
	}
	return domainOrders
}
