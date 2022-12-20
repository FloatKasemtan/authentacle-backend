package gin

import (
	"github.com/floatkasemtan/authentacle-service/handler"
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/init/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	apiHandler := r.Group("api")
	handler.InitGin(apiHandler)

	r.Run(":" + config.C.PORT)
}
