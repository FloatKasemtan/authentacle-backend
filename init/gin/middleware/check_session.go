package middleware

import (
	"github.com/floatkasemtan/authentacle-service/type/response"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CheckSession = func(ctx *gin.Context) {
	id, _, _, err := util.GetUserInfo(ctx)
	// get id and user agent from token
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:  string(http.StatusBadRequest),
			Error: err.Error(),
		})
	}
	print(id)
	// TODO : Check session in database

	ctx.Next()
}
