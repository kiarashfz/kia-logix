package presenter

import (
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/internal/addresses"
)

type AddressResp struct {
	ID          uint                 `json:"id"`
	State       string               `json:"state"`
	City        string               `json:"city"`
	AddressLine string               `json:"address_line"`
	PostalCode  string               `json:"postal_code"`
	Coordinates *helpers.Coordinates `json:"coordinates"`
}

func DomainCoordinatesToCoordinatesResp(domainCoo *addresses.Coordinates) *helpers.Coordinates {
	if domainCoo != nil {
		return &helpers.Coordinates{
			Lat: domainCoo.Lat,
			Lng: domainCoo.Lng,
		}
	}
	return nil
}

func AddressToCreateAddressResp(domainAddress *addresses.Address) *AddressResp {
	return &AddressResp{
		ID:          domainAddress.ID,
		State:       domainAddress.State,
		City:        domainAddress.City,
		AddressLine: domainAddress.AddressLine,
		PostalCode:  domainAddress.PostalCode,
		Coordinates: DomainCoordinatesToCoordinatesResp(domainAddress.Coordinates),
	}
}
