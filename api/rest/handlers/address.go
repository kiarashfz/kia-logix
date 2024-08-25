package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/api/rest/handlers/presenter"
	"kia-logix/api/rest/handlers/requests"
	"kia-logix/pkg/jwt"
	"kia-logix/service"
)

func CreateAddress(AddressService service.IAddressService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.CreateAddress
		if err := c.BodyParser(&req); err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		err := helpers.ValidateBody(req)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		userClaims := c.Locals(jwt.UserClaimKey).(*jwt.UserClaims)

		newAddr := requests.CreateAddressToDomainAddress(&req)
		createdAddress, err := AddressService.Create(c.Context(), newAddr, userClaims.UserID)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusInternalServerError)
		}
		data := presenter.AddressToCreateAddressResp(createdAddress)
		return helpers.SendResponse(c, fiber.StatusCreated, "Address created successfully.", data)
	}
}
