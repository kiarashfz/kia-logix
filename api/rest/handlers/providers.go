package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/api/rest/handlers/presenter"
	"kia-logix/service"
)

func GetProviders(providerService service.IProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, pageSize := helpers.GetPageAndPageSizeFromQParams(c)
		providers, total, err := providerService.GetAll(c.Context(), page, pageSize)
		if err != nil {
			status := fiber.StatusInternalServerError
			return helpers.SendError(c, err, status)
		}
		data := presenter.BatchProviderToGetProvidersResp(providers)
		return helpers.SendPaginatedResponse(c, fiber.StatusOK, "providers successfully fetched.", page, pageSize, total, data)
	}
}
