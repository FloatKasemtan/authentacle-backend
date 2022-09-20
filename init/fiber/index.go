package fiber

import (
	"github.com/floatkasemtan/authentacle-service/handler"
	"log"
	"time"

	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/init/fiber/middleware"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Initialize() {
	// Initialize http instance
	app = fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	})

	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.SuccessResponse{
			Success: true,
			Message: "You are in calling Authentacle. tooo... tooooo... toooo.....",
			Data:    nil,
		})
	})

	// create group /api
	apiGroup := app.Group("api/")

	apiGroup.Use(middleware.Limiter)
	apiGroup.Use(middleware.Cors)

	handler.Init(apiGroup)

	err := app.Listen(":" + config.C.PORT)
	if err != nil {
		log.Fatal(err.Error())
	}

}
