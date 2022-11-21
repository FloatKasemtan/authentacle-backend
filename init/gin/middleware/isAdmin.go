package middleware

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var IsAdmin = func(ctx *gin.Context) {
	_, level, err := util.GetUserInfo(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:  string(http.StatusBadRequest),
			Error: err.Error(),
		})
	}
	if *level != int8(2) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    string(http.StatusBadRequest),
			Message: "User not allow to access this feature",
		})
	}

	ctx.Next()
}
