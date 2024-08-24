package service

import (
	"context"
	"kia-logix/internal/providers"
)

type IProviderService interface {
	GetAll(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error)
}

type ProviderService struct {
	providerOps providers.IProviderOps
}

func NewProviderService(providerOps providers.IProviderOps) *ProviderService {
	return &ProviderService{providerOps: providerOps}
}

func (p *ProviderService) GetAll(ctx context.Context, page, pageSize uint) ([]providers.Provider, uint, error) {
	return p.providerOps.GetAll(ctx, page, pageSize)
}
