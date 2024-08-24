package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers/helpers"
	"kia-logix/api/rest/handlers/requests"
	"kia-logix/internal/user"
	"kia-logix/service"
	"strings"
)

func RegisterUser(authService service.IAuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserRegister
		if err := c.BodyParser(&req); err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		err := helpers.ValidateBody(req)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		u := requests.UserRegisterToUserDomain(&req)
		_, err = authService.CreateUser(c.Context(), u)
		if err != nil {
			if errors.Is(err, user.ErrInvalidPhone) || errors.Is(err, user.ErrInvalidPassword) {
				return helpers.SendError(c, err, fiber.StatusBadRequest)
			}
			if errors.Is(err, user.ErrPhoneAlreadyExists) {
				return helpers.SendError(c, err, fiber.StatusConflict)
			}

			return helpers.SendError(c, err, fiber.StatusInternalServerError)
		}

		return helpers.SendResponse(c, fiber.StatusCreated, "user registered successfully.", nil)
	}
}

func LoginUser(authService service.IAuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserLogin
		if err := c.BodyParser(&req); err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		err := helpers.ValidateBody(req)
		if err != nil {
			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}

		authToken, err := authService.Login(c.Context(), req.Phone, req.Password)
		if err != nil {

			return helpers.SendError(c, err, fiber.StatusBadRequest)
		}
		return helpers.SendResponse(c, fiber.StatusOK, "you logged in successfully.", authToken)
	}
}

func RefreshToken(authService service.IAuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refToken := c.GetReqHeaders()["Authorization"][0]
		if len(refToken) == 0 {
			return helpers.SendError(c, errors.New("token should be provided"), fiber.StatusBadRequest)
		}
		pureToken := strings.Split(refToken, " ")[1]
		authToken, err := authService.RefreshAuth(c.UserContext(), pureToken)
		if err != nil {
			return helpers.SendError(c, errors.New("token is not valid"), fiber.StatusUnauthorized)
		}
		return helpers.SendResponse(c, fiber.StatusOK, "your token successfully refreshed.", authToken)
	}
}
