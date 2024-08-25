package addresses

import "context"

type IAddressOps interface {
	Create(ctx context.Context, address *Address) (*Address, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) Create(ctx context.Context, address *Address) (*Address, error) {
	return o.repo.Create(ctx, address)
}
