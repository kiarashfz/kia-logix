package storage

import (
	"context"
	"gorm.io/gorm"
	"kia-logix/internal/providers"
	"kia-logix/pkg/adapters/storage/entities"
	"kia-logix/pkg/adapters/storage/mappers"
)

type providerRepo struct {
	db *gorm.DB
}

func NewProviderRepo(db *gorm.DB) providers.Repo {
	return &providerRepo{
		db: db,
	}
}

func (p *providerRepo) GetAll(ctx context.Context, limit, offset uint) ([]providers.Provider, uint, error) {
	query := p.db.WithContext(ctx).Model(&entities.Provider{})

	var totalCount int64

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if offset > 0 {
		query = query.Offset(int(offset))
	}

	if limit > 0 {
		query = query.Limit(int(limit))
	}

	var providersEntities []entities.Provider
	if err := query.Find(&providersEntities).Error; err != nil {
		return nil, 0, err
	}

	return mappers.BatchProviderEntitiesToDomain(providersEntities), uint(totalCount), nil
}
