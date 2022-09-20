package user

import (
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) userService {
	return userService{userRepository: userRepository}
}

func (s userService) SignUp(username string, email string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return "", err
	}

	// Test revert check
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return "", err
	}

	userId, err := s.userRepository.SignUp(username, email, string(hashedPassword))
	// Create the Claims
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.C.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return t, nil
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
