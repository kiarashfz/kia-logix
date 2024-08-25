package storage

import (
	"context"
	"gorm.io/gorm"
	"kia-logix/internal/providers"
	"kia-logix/pkg/adapters/storage/entities"
	"kia-logix/pkg/adapters/storage/mappers"
	internalErrors "kia-logix/pkg/errors"
	"strings"
)

type providerRepo struct {
	db *gorm.DB
}

func NewProviderRepo(db *gorm.DB) providers.Repo {
	return &providerRepo{
		db: db,
	}
}

func (p *providerRepo) Create(ctx context.Context, provider *providers.Provider) (*providers.Provider, error) {
	newProvider := mappers.ProviderDomainToEntity(provider)
	err := p.db.WithContext(ctx).Create(&newProvider).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, internalErrors.DBErrDuplicateKey
		}
		return nil, err
	}
	createdProvider := mappers.ProviderEntityToDomain(*newProvider)
	return &createdProvider, nil
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

func (p *providerRepo) GetReports(ctx context.Context, limit, offset uint) (providers []providers.Provider, total uint, err error) {
	query := p.db.WithContext(ctx).Raw(`SELECT
  provider_id AS id,
  providers.name,
  AVG(COALESCE(DATE_PART('day', delivered_date - pickup_date), 0)) AS avg_delivery_days
FROM
  orders
JOIN
  providers ON orders.provider_id = providers.id
WHERE
  pickup_date IS NOT NULL
  AND delivered_date IS NOT NULL
GROUP BY
  provider_id,
  providers.name
ORDER BY
  avg_delivery_days DESC;`)
	err = query.Scan(&providers).Error

	if err != nil {
		return nil, 0, err
	}
	if providers == nil {
		return nil, 0, nil
	}

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

	return providers, uint(totalCount), nil
}
