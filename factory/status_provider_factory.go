package factory

import (
	"fmt"
	"kia-logix/pkg/adapters/status_provider"
)

func GetStatusProvider(providerName, providerURL string) (status_provider.IStatusProvider, error) {
	if providerName == "podro" {
		return status_provider.NewPodro(providerURL), nil
	}
	return nil, fmt.Errorf("wrong provider type passed")
}
