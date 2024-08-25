package orders

import (
	"context"
	"kia-logix/pkg/validations"
)

type IOrderOps interface {
	Create(ctx context.Context, order *Order) (*Order, error)
	Update(ctx context.Context, order *Order) error
	GetUpdateNeedOrders(ctx context.Context) ([]Order, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) Create(ctx context.Context, order *Order) (*Order, error) {
	err := validations.ValidatePhone(order.Sender.Phone)
	if err != nil {
		return nil, ErrInvalidSenderPhone
	}
	err = validations.ValidatePhone(order.Receiver.Phone)
	if err != nil {
		return nil, ErrInvalidReceiverPhone
	}
	order.Sender.Phone = validations.NormalizePhone(order.Sender.Phone)
	order.Receiver.Phone = validations.NormalizePhone(order.Receiver.Phone)
	return o.repo.Create(ctx, order)
}

func (o *Ops) Update(ctx context.Context, order *Order) error {
	return o.repo.Update(ctx, order)
}

func (o *Ops) GetUpdateNeedOrders(ctx context.Context) ([]Order, error) {
	return o.repo.GetUpdateNeedOrders(ctx)
}
