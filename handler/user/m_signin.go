package user

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h userHandler) SignIn(c *gin.Context) {
	body := new(request.UserLoginRequest)

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

	token, isVerify, secret, url, err := h.userService.SignIn(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Username and Password are not match",
			Error:   err.Error(),
		})
		return
	}
	if *isVerify {
		c.JSON(http.StatusOK, response.NewSuccessResponse("Successfully login", map[string]any{"token": token, "is_verify": isVerify}))
	} else {
		c.JSON(http.StatusOK, response.NewSuccessResponse("Please verify your account before login", map[string]any{"token": token, "is_verify": isVerify, "user_secret": secret, "url": url}))
	}
}
