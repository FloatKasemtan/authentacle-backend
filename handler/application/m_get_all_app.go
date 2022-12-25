package application

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h applicationHandler) GetAllApps(c *gin.Context) {
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
		})
		return
	}

	// Get application of user
	applications, err := h.applicationService.GetAllApps(*id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data: map[string]any{
			"applications": applications,
		},
	})
}
