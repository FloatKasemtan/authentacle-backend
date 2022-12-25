package application

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h applicationHandler) CreateApp(c *gin.Context) {
	// Get user id
	id, _, verified, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		return
	}
	if !*verified {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Session unverified",
			Error:   err.Error(),
		})
		return
	}

	// Parse request body
	body := new(request.ApplicationRequest)
	if err := c.ShouldBindBodyWith(body, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if err := validator.Validate.Struct(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if err := h.applicationService.CreateApp(body, *id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Successfully create application",
	})

}
