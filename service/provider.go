package service

import (
	"context"
	"kia-logix/internal/providers"
)

type IProviderService interface {
	Create(ctx context.Context, provider *providers.Provider) (*providers.Provider, error)
	GetAll(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error)
	GetProvidersReports(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error)
}

type ProviderService struct {
	providerOps providers.IProviderOps
}

func NewProviderService(providerOps providers.IProviderOps) *ProviderService {
	return &ProviderService{providerOps: providerOps}
}

func (p *ProviderService) Create(ctx context.Context, provider *providers.Provider) (*providers.Provider, error) {
	return p.providerOps.Create(ctx, provider)
}

func (p *ProviderService) GetAll(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error) {
	return p.providerOps.GetAll(ctx, page, pageSize)
}

func (p *ProviderService) GetProvidersReports(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error) {
	return p.providerOps.GetReports(ctx, page, pageSize)
}
