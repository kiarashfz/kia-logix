package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers"
	"kia-logix/api/rest/middlewares"
	"kia-logix/api/rest/permissions"
	"kia-logix/cmd/app"
	"kia-logix/config"
	"log"
)

func Run(cfg config.Config, appContainer *app.Container) {
	fiberApp := fiber.New()
	fiberApp.Use(middlewares.LimiterMiddleware(1, 100, cfg.Redis))
	apiV1 := fiberApp.Group("/api/v1")

	registerGlobalRoutes(apiV1, appContainer)
	registerProviderRoutes(apiV1, appContainer, []byte(cfg.Auth.SecretToken))
	registerAddressRoutes(apiV1, appContainer, []byte(cfg.Auth.SecretToken))
	registerOrderRoutes(apiV1, appContainer, []byte(cfg.Auth.SecretToken))

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerGlobalRoutes(router fiber.Router, app *app.Container) {
	router.Post("/register", handlers.RegisterUser(app.GetAuthService()))
	router.Post("/login", handlers.LoginUser(app.GetAuthService()))
	router.Get("/refresh", handlers.RefreshToken(app.GetAuthService()))
}

func registerProviderRoutes(router fiber.Router, app *app.Container, secret []byte) {
	router = router.Group("/providers")
	router.Use(middlewares.Auth(secret))
	router.Use(permissions.AuthenticatedPermission())
	router.Use(permissions.AdminPermission())

	router.Post("",
		handlers.CreateProvider(app.GetProviderService()),
	)
	router.Get("",
		handlers.GetProviders(app.GetProviderService()),
	)
}

func registerAddressRoutes(router fiber.Router, app *app.Container, secret []byte) {
	router = router.Group("/addresses")
	router.Use(middlewares.Auth(secret))
	router.Use(permissions.AuthenticatedPermission())

	router.Post("",
		handlers.CreateAddress(app.GetAddressService()),
	)
}

func registerOrderRoutes(router fiber.Router, app *app.Container, secret []byte) {
	router = router.Group("/orders")
	router.Use(middlewares.Auth(secret))
	router.Use(permissions.AuthenticatedPermission())

	router.Post("",
		handlers.CreateOrder(app.GetOrderService()),
	)
}
