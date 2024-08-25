package service

import (
	"context"
	worker "github.com/contribsys/faktory_worker_go"
	"kia-logix/service/mappers"
)

func (o *OrderService) UpdateOrderStatusTask(ctx context.Context, args ...interface{}) error {
	worker.HelperFor(ctx)
	orderMap := args[0].(map[string]interface{})
	domainOrder := mappers.OrderArgsMapToDomain(orderMap)
	err := o.UpdateOrderStatusJob(ctx, domainOrder)
	if err != nil {
		return err
	}
	return nil
}
