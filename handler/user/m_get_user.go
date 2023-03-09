package user

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    "",
	})

	id, _, _, err := util.GetUserInfo(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		return
	}

	user, err := h.userService.GetUser(*id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Unable to get user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    user,
	})
}
