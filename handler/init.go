package handler

import (
	"github.com/floatkasemtan/authentacle-service/handler/application"
	"github.com/floatkasemtan/authentacle-service/handler/user"
	"github.com/floatkasemtan/authentacle-service/init/db"
	"github.com/floatkasemtan/authentacle-service/init/gin/middleware"
	appRepo "github.com/floatkasemtan/authentacle-service/repository/application"
	userRepo "github.com/floatkasemtan/authentacle-service/repository/user"
	appService "github.com/floatkasemtan/authentacle-service/service/application"
	userService "github.com/floatkasemtan/authentacle-service/service/user"
	"github.com/gin-gonic/gin"
)

func InitGin(router *gin.RouterGroup) {
	// User Endpoints
	userRepository := userRepo.NewUserRepositoryDB(db.DB)
	userService := userService.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	userGroup := router.Group("user")

	userGroup.POST("login", userHandler.SignIn)
	userGroup.POST("register", userHandler.SignUp)
	userGroup.POST("check-otp", userHandler.CheckOTP)
	userGroup.PUT("verify", userHandler.Verify)

	// Application Endpoints
	applicationRepository := appRepo.NewAppRepositoryDB(db.DB)
	applicationService := appService.NewAppService(applicationRepository)
	applicationHandler := application.NewAppHandler(applicationService)

	applicationGroup := router.Group("application")

	applicationGroup.GET("all", applicationHandler.GetAllApps)
	applicationGroup.GET(":id", applicationHandler.GetApp)
	applicationGroup.POST("create", applicationHandler.CreateApp)

	// Administrator Endpoints
	adminGroup := router.Group("admin", middleware.IsAdmin)

	adminGroup.GET("get", userHandler.GetUser)
}
