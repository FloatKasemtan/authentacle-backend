package user

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h userHandler) CheckOTP(c *gin.Context) {
	id, role, _, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		return
	}

	otp := c.Query("otp")
	token, err := h.userService.CheckOTP(*id, *role, otp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Unable to verify",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    token,
	})
}
