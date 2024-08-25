package requests

import (
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/internal/addresses"
)

type CreateAddress struct {
	State       string               `json:"state" validate:"required"`
	City        string               `json:"city" validate:"required"`
	AddressLine string               `json:"address_line" validate:"required"`
	PostalCode  string               `json:"postal_code"`
	Coordinates *helpers.Coordinates `json:"coordinates"`
}

func CreateAddressToDomainAddress(a *CreateAddress) *addresses.Address {
	var domainCoordinates *addresses.Coordinates
	if a.Coordinates != nil {
		domainCoordinates = &addresses.Coordinates{
			Lat: a.Coordinates.Lat,
			Lng: a.Coordinates.Lng,
		}
	}
	return &addresses.Address{
		State:       a.State,
		City:        a.City,
		AddressLine: a.AddressLine,
		PostalCode:  a.PostalCode,
		Coordinates: domainCoordinates,
	}
}
