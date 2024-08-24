package providers

import (
	"context"
)

type IProviderOps interface {
	GetAll(ctx context.Context, page, pageSize uint) ([]Provider, uint, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) GetAll(ctx context.Context, page, pageSize uint) ([]Provider, uint, error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	return o.repo.GetAll(ctx, limit, offset)
}
