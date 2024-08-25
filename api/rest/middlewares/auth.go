package middlewares

import (
	"errors"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/pkg/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(secret []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return helpers.SendError(c, errors.New("authorization header missing"), fiber.StatusUnauthorized)
		}

		// Split the Authorization header value
		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return helpers.SendError(c, errors.New("invalid authorization token format"), fiber.StatusUnauthorized)
		}

		//pureToken := parts[1]
		pureToken := parts[1]
		claims, err := jwt.ParseToken(pureToken, secret)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusUnauthorized)
		}

		c.Locals(jwt.UserClaimKey, claims)

		return c.Next()
	}
}
