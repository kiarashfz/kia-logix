package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/api/rest/handlers/presenter"
	"kia-logix/api/rest/handlers/requests"
	"kia-logix/internal/providers"
	"kia-logix/service"
)

func CreateProvider(providerService service.IProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.CreateProvider
		if err := c.BodyParser(&req); err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		err := helpers.ValidateBody(req)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		newProvider := requests.CreateProviderToDomainProvider(&req)
		createdProvider, err := providerService.Create(c.Context(), newProvider)
		if err != nil {
			if errors.Is(err, providers.ErrProviderURLInvalid) {
				return helpers.SendError(c, err, fiber.StatusBadRequest)
			}
			if errors.Is(err, providers.ErrProviderNameAlreadyExists) {
				return helpers.SendError(c, err, fiber.StatusConflict)
			}
			return helpers.SendError(c, err, fiber.StatusInternalServerError)
		}
		data := presenter.ProviderToProviderResp(*createdProvider)
		return helpers.SendResponse(c, fiber.StatusCreated, "provider created successfully.", data)
	}
}

func GetProviders(providerService service.IProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, pageSize := helpers.GetPageAndPageSizeFromQParams(c)
		fetchedProviders, total, err := providerService.GetAll(c.Context(), page, pageSize)
		if err != nil {
			status := fiber.StatusInternalServerError
			return helpers.SendError(c, err, status)
		}
		data := presenter.BatchProviderToGetProvidersResp(fetchedProviders)
		return helpers.SendPaginatedResponse(c, fiber.StatusOK, "providers successfully fetched.", page, pageSize, total, data)
	}
}
