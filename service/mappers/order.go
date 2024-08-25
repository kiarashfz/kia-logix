package mappers

import (
	"kia-logix/internal/orders"
	"time"
)

func OrderArgsMapToDomain(orderMap map[string]interface{}) *orders.Order {
	var pickupDatePtr *time.Time
	var deliveredDatePtr *time.Time
	if orderMap["PickupDate"] != nil {
		pickupDate, _ := time.Parse("2006-01-02", orderMap["PickupDate"].(string))
		pickupDatePtr = &pickupDate
	}
	if orderMap["DeliveredDate"] != nil {
		deliveredDate, _ := time.Parse("2006-01-02", orderMap["DeliveredDate"].(string))
		deliveredDatePtr = &deliveredDate
	}
	return &orders.Order{
		ID:                uint(orderMap["ID"].(float64)),
		OwnerID:           uint(orderMap["OwnerID"].(float64)),
		Owner:             UserArgsMapToDomain(orderMap["Owner"].(map[string]interface{})),
		ProviderID:        uint(orderMap["ProviderID"].(float64)),
		Provider:          ProviderArgsMapToDomain(orderMap["Provider"].(map[string]interface{})),
		PickupDate:        pickupDatePtr,
		DeliveredDate:     deliveredDatePtr,
		Status:            orders.OrderStatus(orderMap["Status"].(string)),
		IsPickedUpSMSSent: orderMap["IsPickedUpSMSSent"].(bool),
	}
}
