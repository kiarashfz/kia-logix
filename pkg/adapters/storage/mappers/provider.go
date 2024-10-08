package mappers

import (
	"kia-logix/internal/providers"
	"kia-logix/pkg/adapters/storage/entities"
	"kia-logix/pkg/fp"
)

func ProviderDomainToEntity(domainProvider *providers.Provider) *entities.Provider {
	return &entities.Provider{
		Name: domainProvider.Name,
		URL:  domainProvider.URL,
	}
}

func ProviderEntityToDomain(providerEntity entities.Provider) providers.Provider {
	return providers.Provider{
		ID:   providerEntity.ID,
		Name: providerEntity.Name,
		URL:  providerEntity.URL,
	}
}

func BatchProviderEntitiesToDomain(providerEntities []entities.Provider) []providers.Provider {
	return fp.Map(providerEntities, ProviderEntityToDomain)
}
