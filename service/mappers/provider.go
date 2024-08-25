package mappers

import "kia-logix/internal/providers"

func ProviderArgsMapToDomain(providerMap map[string]interface{}) *providers.Provider {
	return &providers.Provider{
		ID:   uint(providerMap["ID"].(float64)),
		Name: providerMap["Name"].(string),
		URL:  providerMap["URL"].(string),
	}
}
