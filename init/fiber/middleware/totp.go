package middleware

import (
	"errors"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gofiber/fiber/v2"
)

var Totp = func(ctx *fiber.Ctx) error {
	if len(ctx.Query("otp")) != 6 {
		return &response.GenericError{
			Code: "400",
			Err:  errors.New("Invalid OTP"),
		}
	}

	return ctx.Next()
}
