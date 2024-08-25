package requests

import (
	"kia-logix/internal/providers"
)

type CreateProvider struct {
	Name string `json:"name" validate:"required"`
	URL  string `json:"url" validate:"required"`
}

func CreateProviderToDomainProvider(p *CreateProvider) *providers.Provider {
	return &providers.Provider{
		Name: p.Name,
		URL:  p.URL,
	}
}
