package user

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h userHandler) SignUp(c *gin.Context) {
	// Parse request body
	body := new(request.UserRequest)
	if err := c.ShouldBindBodyWith(body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}
	if err := validator.Validate.Struct(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	// Create user
	token, url, secret, err := h.userService.SignUp(body.Username, body.Email, body.Password, c.Request.UserAgent())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Successfully register",
		Data: map[string]any{
			"token":       token,
			"url":         url,
			"user_secret": secret,
		},
	})
}
