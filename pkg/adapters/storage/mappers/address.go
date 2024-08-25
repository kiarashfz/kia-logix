package mappers

import (
	"kia-logix/internal/addresses"
	"kia-logix/pkg/adapters/storage/entities"
)

func AddressDomainToEntity(domainAddress *addresses.Address) *entities.Address {
	return &entities.Address{
		State:       domainAddress.State,
		City:        domainAddress.City,
		PostalCode:  domainAddress.PostalCode,
		AddressLine: domainAddress.AddressLine,
		Coordinates: domainAddress.Coordinates,
		OwnerID:     domainAddress.OwnerID,
	}
}

func AddressEntityToDomain(addressEntity *entities.Address) *addresses.Address {
	return &addresses.Address{
		ID:          addressEntity.ID,
		State:       addressEntity.State,
		City:        addressEntity.City,
		AddressLine: addressEntity.AddressLine,
		PostalCode:  addressEntity.PostalCode,
		Coordinates: addressEntity.Coordinates,
	}
}
