package helpers

import (
	"github.com/gofiber/fiber/v2"
	"kia-logix/pkg/valuecontext"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func SendError(c *fiber.Ctx, err error, status int) error {
	c.Locals(valuecontext.IsTxError, err)

	return c.Status(status).JSON(map[string]any{
		"error_msg": err.Error(),
	})
}

func SendResponse(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(Response{
		Message: message,
		Data:    data,
	})
}
