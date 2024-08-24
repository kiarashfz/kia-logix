package presenter

import (
	"kia-logix/internal/providers"
	"kia-logix/pkg/fp"
)

type ProviderResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func providerToProviderResp(domainProvider providers.Provider) ProviderResp {
	return ProviderResp{
		ID:   domainProvider.ID,
		Name: domainProvider.Name,
		URL:  domainProvider.URL,
	}
}

func BatchProviderToGetProvidersResp(domainProviders []providers.Provider) []ProviderResp {
	return fp.Map(domainProviders, providerToProviderResp)
}
