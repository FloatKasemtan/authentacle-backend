package application

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/service/application"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type applicationHandler struct {
	applicationService application.ApplicationService
}

func NewAppHandler(applicationService application.ApplicationService) applicationHandler {
	return applicationHandler{applicationService: applicationService}
}

func (h applicationHandler) GetAllApps(c *gin.Context) {
	// Get user id
	id, _, verified, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
	}
	if !*verified {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Session unverified",
			Error:   err.Error(),
		})
	}

	// Get application of user
	applications, err := h.applicationService.GetAllApps(*id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data: map[string]any{
			"applications": applications,
		},
	})
}

func (h applicationHandler) GetApp(c *gin.Context) {
	// Get user id
	id, _, verified, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
	}
	if !*verified {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Session unverified",
			Error:   err.Error(),
		})
	}

	appId := c.Params.ByName("id")

	// Parse request body
	app, err := h.applicationService.GetApp(appId, *id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data: map[string]any{
			"application": app,
		},
	})

}

func (h applicationHandler) CreateApp(c *gin.Context) {
	// Get user id
	id, _, verified, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
	}
	if !*verified {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Session unverified",
			Error:   err.Error(),
		})
	}

	// Parse request body
	body := new(request.ApplicationRequest)
	if err := c.ShouldBindBodyWith(body, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validator.Validate.Struct(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := h.applicationService.CreateApp(body, *id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Successfully create application",
	})

}
