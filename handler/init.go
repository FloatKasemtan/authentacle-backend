package handler

import (
	"github.com/floatkasemtan/authentacle-service/handler/application"
	"github.com/floatkasemtan/authentacle-service/handler/user"
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/init/db"
	"github.com/floatkasemtan/authentacle-service/init/fiber/middleware"
	appRepo "github.com/floatkasemtan/authentacle-service/repository/application"
	userRepo "github.com/floatkasemtan/authentacle-service/repository/user"
	appService "github.com/floatkasemtan/authentacle-service/service/application"
	userService "github.com/floatkasemtan/authentacle-service/service/user"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Init(router fiber.Router) {
	// User Endpoints
	userRepository := userRepo.NewUserRepositoryDB(db.DB)
	userService := userService.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	userGroup := router.Group("user/")

	userGroup.Post("login", middleware.Totp, userHandler.SignIn)
	userGroup.Post("register", userHandler.SignUp)

	applicationRepository := appRepo.NewAppRepositoryDB(db.DB)
	applicationService := appService.NewAppService(applicationRepository)
	applicationHandler := application.NewAppHandler(applicationService)

	applicationGroup := router.Group("application/")

	applicationGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.C.JWT_SECRET),
	}))

	applicationGroup.Get("all", applicationHandler.GetAllApps)
	applicationGroup.Get(":id", middleware.Totp, applicationHandler.GetApp)
	applicationGroup.Post("create", applicationHandler.CreateApp)
}
