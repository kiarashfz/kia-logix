package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kia-logix/cmd/api/app"
	"kia-logix/config"
	"log"
)

func Run(cfg config.Config, app *app.Container) {
	fiberApp := fiber.New()
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}
