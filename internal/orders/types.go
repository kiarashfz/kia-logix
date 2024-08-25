package orders

import (
	"context"
	"kia-logix/internal/addresses"
	"kia-logix/internal/providers"
	"kia-logix/internal/user"
	"time"
)

type Repo interface {
	Create(ctx context.Context, order *Order) (*Order, error)
}

type Sender struct {
	Name      string
	Phone     string
	AddressID uint
	Address   *addresses.Address
}

type Receiver struct {
	Name      string
	Phone     string
	AddressID uint
	Address   *addresses.Address
}

type Order struct {
	ID         uint
	OwnerID    uint
	Owner      *user.User
	Sender     *Sender
	Receiver   *Receiver
	ProviderID uint
	Provider   *providers.Provider
	PickupDate time.Time
	Status     OrderStatus
}

type OrderStatus string

const (
	REGISTERED    OrderStatus = "REGISTERED"
	PICKED_UP     OrderStatus = "PICKED_UP"
	PROVIDER_SEEN OrderStatus = "PROVIDER_SEEN"
	IN_PROGRESS   OrderStatus = "IN_PROGRESS"
	PENDING       OrderStatus = "PENDING"
	DELIVERED     OrderStatus = "DELIVERED"
)
