package application

import (
	"github.com/floatkasemtan/authentacle-service/service/application"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type applicationHandler struct {
	applicationService application.ApplicationService
}

func NewAppHandler(applicationService application.ApplicationService) applicationHandler {
	return applicationHandler{applicationService: applicationService}
}

func (h applicationHandler) GetAllApps(c *fiber.Ctx) error {
	// Get user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// Get application of user
	applications, err := h.applicationService.GetAllApps(id)
	if err != nil {
		return c.JSON(response.ErrorResponse{
			Code:  "400",
			Error: err.Error(),
		})
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data: map[string]any{
			"applications": applications,
		},
	})
}

func (h applicationHandler) GetApp(c *fiber.Ctx) error {
	// Get user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// Parse request body

	app, err := h.applicationService.GetApp(c.Params("id"), id)
	if err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data: map[string]any{
			"application": app,
		},
	})
}

func (h applicationHandler) CreateApp(c *fiber.Ctx) error {
	// Get user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// Parse request body
	body := new(request.ApplicationRequest)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}

	if err := h.applicationService.CreateApp(body, id); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
	})
}
