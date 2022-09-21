package user

import (
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pquerna/otp/totp"
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

	userId, err := s.userRepository.SignUp(username, email, string(hashedPassword))
	if err != nil {
		return "", err
	}
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
	user, err := s.userRepository.SignIn(username)
	if err != nil {
		return "", err
	}

	// Test revert check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  user.Id,
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

func (s userService) GetUser(userId string) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s userService) SendVerificationForm(id string, email string) error {
	generate, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "authentacle.floatykt.com",
		AccountName: email,
	})

	if err != nil {
		return err
	}
	if err := s.userRepository.SendVerificationForm(id, generate.Secret()); err != nil {
		return err
	}
	return nil
}
