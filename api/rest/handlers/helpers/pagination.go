package helpers

import (
	"github.com/gofiber/fiber/v2"
	"math"
)

type PaginatedResponse[T any] struct {
	Message    string `json:"message,omitempty"`
	Page       uint   `json:"page"`
	PageSize   uint   `json:"page_size"`
	TotalPages uint   `json:"total_pages"`
	Data       []T    `json:"data"`
}

func GetPageAndPageSizeFromQParams(c *fiber.Ctx) (uint, uint) {
	page, pageSize := c.QueryInt("page"), c.QueryInt("page_size")
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	return uint(page), uint(pageSize)
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
