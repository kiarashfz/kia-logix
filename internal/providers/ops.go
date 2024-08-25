package providers

import (
	"context"
	"errors"
	internalErrors "kia-logix/pkg/errors"
)

type IProviderOps interface {
	Create(ctx context.Context, provider *Provider) (*Provider, error)
	GetAll(ctx context.Context, page, pageSize uint) ([]Provider, uint, error)
	GetReports(ctx context.Context, page, pageSize uint) ([]Provider, uint, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) Create(ctx context.Context, provider *Provider) (*Provider, error) {
	err := ValidateURL(provider.URL)
	if err != nil {
		return nil, ErrProviderURLInvalid
	}
	createdProvider, err := o.repo.Create(ctx, provider)
	if err != nil {
		if errors.Is(err, internalErrors.DBErrDuplicateKey) {
			return nil, ErrProviderNameAlreadyExists
		}
		return nil, err
	}
	return createdProvider, nil
}

func (o *Ops) GetAll(ctx context.Context, page, pageSize uint) ([]Provider, uint, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	return o.repo.GetAll(ctx, limit, offset)
}

func (o *Ops) GetReports(ctx context.Context, page, pageSize uint) ([]Provider, uint, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	return o.repo.GetReports(ctx, limit, offset)
}
