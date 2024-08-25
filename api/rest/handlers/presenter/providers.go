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

type ProviderReportResp struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	AvgDeliveryDays float64 `json:"avg_delivery_days"`
}

func ProviderToProviderResp(domainProvider providers.Provider) ProviderResp {
	return ProviderResp{
		ID:   domainProvider.ID,
		Name: domainProvider.Name,
		URL:  domainProvider.URL,
	}
}

func ProviderToProviderReportResp(domainProvider providers.Provider) ProviderReportResp {
	return ProviderReportResp{
		ID:              domainProvider.ID,
		Name:            domainProvider.Name,
		AvgDeliveryDays: domainProvider.AvgDeliveryDays,
	}
}

func BatchProviderToGetProvidersResp(domainProviders []providers.Provider) []ProviderResp {
	return fp.Map(domainProviders, ProviderToProviderResp)
}

func BatchProviderToGetProvidersReportsResp(domainProviders []providers.Provider) []ProviderReportResp {
	return fp.Map(domainProviders, ProviderToProviderReportResp)
}
