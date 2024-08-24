package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kia-logix/api/rest/handlers"
	"kia-logix/api/rest/middlewares"
	"kia-logix/cmd/api/app"
	"kia-logix/config"
	"log"
)

func Run(cfg config.Config, appContainer *app.Container) {
	fiberApp := fiber.New()
	fiberApp.Use(middlewares.LimiterMiddleware(1, 100, cfg.Redis))
	apiV1 := fiberApp.Group("/api/v1")

	registerProviderRoutes(apiV1, appContainer)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerProviderRoutes(router fiber.Router, app *app.Container) {
	router = router.Group("/providers")

	router.Get("",
		handlers.GetProviders(app.ProviderService()),
	)
}
