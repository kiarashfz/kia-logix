package service

import (
	"context"
	"kia-logix/internal/addresses"
)

type IAddressService interface {
	Create(ctx context.Context, address *addresses.Address, userID uint) (*addresses.Address, error)
}

type AddressService struct {
	addressOps addresses.IAddressOps
}

func NewAddressService(addressOps addresses.IAddressOps) *AddressService {
	return &AddressService{addressOps: addressOps}
}

func (o *AddressService) Create(ctx context.Context, address *addresses.Address, userID uint) (*addresses.Address, error) {
	address.OwnerID = userID
	return o.addressOps.Create(ctx, address)
}
