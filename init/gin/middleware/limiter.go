package middleware

import (
	"time"

	"github.com/floatkasemtan/authentacle-service/type/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = func() fiber.Handler {
	config := limiter.Config{
		Next:       nil,
		Expiration: 60 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return &response.GenericError{
				Message: "Rate limit reached, try again in a few minutes.",
			}
		},
	}

	return limiter.New(config)
}()
