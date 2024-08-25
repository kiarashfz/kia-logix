package orders

import (
	"context"
)

type IOrderOps interface {
	Create(ctx context.Context, order *Order) (*Order, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) Create(ctx context.Context, order *Order) (*Order, error) {
	return o.repo.Create(ctx, order)
}
