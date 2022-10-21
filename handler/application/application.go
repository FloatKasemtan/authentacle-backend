package application

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/service/application"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gofiber/fiber/v2"
)

type applicationHandler struct {
	applicationService application.ApplicationService
}

func NewAppHandler(applicationService application.ApplicationService) applicationHandler {
	return applicationHandler{applicationService: applicationService}
}

func (h applicationHandler) GetAllApps(c *fiber.Ctx) error {
	// Get user id
	id := util.GetUserId(c)

	// Get application of user
	applications, err := h.applicationService.GetAllApps(id)
	if err != nil {
		return err
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
	id := util.GetUserId(c)

	// Parse request body
	app, err := h.applicationService.GetApp(c.Params("id"), id)
	if err != nil {
		return err
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
	id := util.GetUserId(c)

	// Parse request body
	body := new(request.ApplicationRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := validator.Validate.Struct(body)
	if err != nil {
		return err
	}

	if err := h.applicationService.CreateApp(body, id); err != nil {
		return err
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "",
	})
}
