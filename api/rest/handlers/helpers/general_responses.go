package helpers

import (
	"github.com/gofiber/fiber/v2"
	"kia-logix/pkg/valuecontext"
	"math"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type PaginatedResponse[T any] struct {
	Message    string `json:"message,omitempty"`
	Page       uint   `json:"page"`
	PageSize   uint   `json:"page_size"`
	TotalPages uint   `json:"total_pages"`
	Data       []T    `json:"data"`
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

func SendPaginatedResponse[T any](c *fiber.Ctx, status int, message string, page, pageSize, total uint, data []T) error {
	totalPages := uint(0)
	if pageSize > 0 && total > 0 {
		totalPages = uint(math.Ceil(float64(total) / float64(pageSize)))
	}
	return c.Status(status).JSON(PaginatedResponse[T]{
		Message:    message,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		Data:       data,
	})
}
