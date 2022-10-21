package middleware

import (
	"errors"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gofiber/fiber/v2"
)

var IsAdmin = func(ctx *fiber.Ctx) error {
	level := util.GetUserAuthorization(ctx)
	if level != 2 {
		return errors.New("user not authorize")
	}

	return ctx.Next()
}
