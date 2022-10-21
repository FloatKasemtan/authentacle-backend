package user

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/service/user"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) SignUp(c *fiber.Ctx) error {
	// Parse request body
	u := new(request.UserRequest)
	if err := c.BodyParser(u); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}

	err := validator.Validate.Struct(u)
	if err != nil {
		return err
	}

	// Create user
	token, base64, secret, err := h.userService.SignUp(u.Username, u.Email, u.Password)
	if err != nil {
		return c.JSON(response.ErrorResponse{Code: "500", Message: err.Error()})
	}

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "Successfully register",
		Data: map[string]any{
			"token":       token,
			"image":       base64,
			"user_secret": secret,
		},
	})
}

func (h userHandler) SignIn(c *fiber.Ctx) error {
	u := new(request.UserLoginRequest)
	if err := c.BodyParser(u); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}

	err := validator.Validate.Struct(u)
	if err != nil {
		return err
	}

	token, err := h.userService.SignIn(u.Username, u.Password, u.Otp)
	if err != nil {
		return c.JSON(response.ErrorResponse{
			Code:    "400",
			Message: "Username and Password are not match",
			Error:   err.Error(),
		})
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data: map[string]any{
			"token": token,
		},
	})
}

func (h userHandler) GetUser(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}

func (h userHandler) Verify(c *fiber.Ctx) error {
	id := util.GetUserId(c)
	otp := c.Query("otp")
	h.userService.Verify(id, otp)
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}
