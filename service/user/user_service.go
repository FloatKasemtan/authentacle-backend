package user

import (
	"fmt"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) userService {
	return userService{userRepository: userRepository}
}

func (s userService) SignUp(username string, email string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	fmt.Printf("Hashed password: %s", string(hashedPassword))

	// Test revert check
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return "", err
	}

	token, err := s.userRepository.SignUp(username, email, string(hashedPassword))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s userService) SignIn(username string, password string) (string, error) {
	token, err := s.userRepository.SignIn(username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s userService) GetUser(userId string) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s userService) SendVerificationForm(id string) error {
	//TODO implement me
	panic("implement me")
}
