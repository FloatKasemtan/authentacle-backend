package user

import (
	"github.com/floatkasemtan/authentacle-service/service/user"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
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

	// Create user
	token, err := h.userService.SignUp(u.Username, u.Email, u.Password)
	if err != nil {
		return c.JSON(response.ErrorResponse{Code: "500", Message: err.Error()})
	}

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "Successfully register",
		Data: map[string]any{
			"token": token,
		},
	})
}

func (h userHandler) SignIn(c *fiber.Ctx) error {
	u := new(request.UserLoginRequest)
	if err := c.BodyParser(u); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}

	token, err := h.userService.SignIn(u.Username, u.Password)
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
			token: "token",
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

func (h userHandler) SendVerificationForm(c *fiber.Ctx) error {

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}
