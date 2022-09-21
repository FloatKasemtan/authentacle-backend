package middleware

import (
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	origins := ""
	for i, s := range config.C.CORS {
		origins += s
		if i < len(config.C.CORS)-1 {
			origins += ", "
		}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}()
