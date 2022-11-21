package middleware

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Error:   err.Err.Error(),
		})
	}

}
