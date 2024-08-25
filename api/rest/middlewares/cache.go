package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/storage/redis/v3"
	"kia-logix/config"
	"runtime"
	"time"
)

// CacheMiddleware Define a function to configure cache middleware
func CacheMiddleware(expMinutes int, cfg config.Redis) fiber.Handler {
	exp := time.Duration(expMinutes)
	return cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   exp * time.Minute,
		CacheControl: true,
		Storage: redis.New(redis.Config{
			Host:      cfg.Host,
			Port:      cfg.Port,
			Password:  cfg.Pass,
			Database:  1,
			Reset:     false,
			TLSConfig: nil,
			PoolSize:  10 * runtime.GOMAXPROCS(0),
		}),
	})
}
