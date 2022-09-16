package application

import (
	"github.com/davecgh/go-spew/spew"
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
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}

func (h applicationHandler) GetApp(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}

func (h applicationHandler) CreateApp(c *fiber.Ctx) error {
	// Get user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	spew.Dump(claims)
	id := claims["id"].(string)

	// Parse request body
	app := new(request.ApplicationRequest)
	if err := c.BodyParser(app); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}
	spew.Dump(app)

	if err := h.applicationService.CreateApp(app, id); err != nil {
		return c.JSON(response.ErrorResponse{Code: "400", Message: err.Error()})
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    nil,
	})
}
