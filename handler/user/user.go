package user

import (
	"github.com/floatkasemtan/authentacle-service/init/validator"
	"github.com/floatkasemtan/authentacle-service/service/user"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) userHandler {
	return userHandler{userService: userService}
}

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
	token, url, secret, err := h.userService.SignUp(body.Username, body.Email, body.Password)
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
		c.JSON(http.StatusOK, response.SuccessResponse{
			Success: true,
			Data:    map[string]any{"token": token, "is_verify": isVerify},
		})
	} else {
		c.JSON(http.StatusOK, response.SuccessResponse{
			Success: true,
			Data:    map[string]any{"token": token, "is_verify": isVerify, "url": url, "user_secret": secret},
		})
	}
}

func (h userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    nil,
	})
}

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

func (h userHandler) Verify(c *gin.Context) {
	id, role, _, err := util.GetUserInfo(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		return
	}

	otp := c.Query("otp")
	token, err := h.userService.Verify(*id, *role, otp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Unable to verify",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    token,
	})
}
