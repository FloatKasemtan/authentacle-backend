package user

import (
	"github.com/floatkasemtan/authentacle-service/service/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) userHandler {
	return userHandler{userService: userService}
}
