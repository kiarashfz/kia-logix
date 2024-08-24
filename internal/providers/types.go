package providers

import (
	"context"
)

type Repo interface {
	GetAll(ctx context.Context, limit, offset uint) (providers []Provider, total uint, err error)
}

type Provider struct {
	ID   uint
	Name string
	URL  string
}
