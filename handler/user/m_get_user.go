package user

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    nil,
	})
}
