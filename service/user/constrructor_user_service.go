package user

import (
	"github.com/floatkasemtan/authentacle-service/repository/user"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) userService {
	return userService{userRepository: userRepository}
}
