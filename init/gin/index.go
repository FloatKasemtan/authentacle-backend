package gin

import (
	"github.com/floatkasemtan/authentacle-service/handler"
	"github.com/floatkasemtan/authentacle-service/init/gin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	apiHandler := r.Group("api")
	handler.InitGin(apiHandler)
}
