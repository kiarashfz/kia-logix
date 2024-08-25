package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/api/rest/handlers/presenter"
	"kia-logix/api/rest/handlers/requests"
	"kia-logix/internal/orders"
	"kia-logix/pkg/jwt"
	"kia-logix/service"
)

func CreateOrder(orderService service.IOrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.CreateOrder
		if err := c.BodyParser(&req); err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		err := helpers.ValidateBody(req)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		userClaims := c.Locals(jwt.UserClaimKey).(*jwt.UserClaims)

		newOrder := requests.CreateOrderToDomainOrder(&req)
		createdOrder, err := orderService.Create(c.Context(), newOrder, userClaims.UserID)
		if err != nil {
			if errors.Is(err, orders.ErrInvalidSenderPhone) || errors.Is(err, orders.ErrInvalidReceiverPhone) {
				return helpers.SendError(c, err, fiber.StatusBadRequest)
			}
			return helpers.SendError(c, err, fiber.StatusInternalServerError)
		}
		data := presenter.OrderToCreateOrderResp(createdOrder)
		return helpers.SendResponse(c, fiber.StatusCreated, "order created successfully.", data)
	}
}
