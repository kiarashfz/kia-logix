package permissions

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/pkg/jwt"
)

func AuthenticatedPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := c.Locals(jwt.UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return helpers.SendError(c, errors.New("wrong claim type"), fiber.StatusBadRequest)
		}
		return c.Next()
	}
}

func AdminPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims := c.Locals(jwt.UserClaimKey).(*jwt.UserClaims)
		if userClaims.IsAdmin == false {
			return helpers.SendError(c, errors.New("you don't have permission"), fiber.StatusForbidden)
		}

		return c.Next()
	}
}
