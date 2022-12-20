package user

import (
	"bytes"
	"errors"
	"github.com/floatkasemtan/authentacle-service/init"
	"image/png"
	"time"

	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) userService {
	return userService{userRepository: userRepository}
}

func (s userService) SignUp(username string, email string, password string) (*string, *string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return nil, nil, err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Authentacle",
		AccountName: email,
	})

	userId, err := s.userRepository.SignUp(username, email, string(hashedPassword), key.Secret())
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, nil, err
	}
	if err := png.Encode(&buf, img); err != nil {
		return nil, nil, err
	}
	secret := key.Secret()
	return &t, &secret, nil
}

func (s userService) SignIn(username string, password string, otp string) (*string, error) {
	user, err := s.userRepository.SignIn(username)
	if err != nil {
		return nil, err
	}

	if !totp.Validate(otp, user.Secret) {
		return nil, errors.New("Invalid OTP")
	}

	// Test revert check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
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
		return nil, err
	}

	return &t, nil
}

func (s userService) GetUser(userId string) (*Response, error) {
	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return nil, err
	}

	return &Response{
		Username: user.Username,
		Email:    user.Email,
		IsVerify: user.IsVerify,
	}, nil
}

func (s userService) Verify(id string, role int8, otp string) (*string, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if !totp.Validate(otp, user.Secret) {
		return nil, errors.New("Invalid OTP")
	}

	err = s.userRepository.Verify(id)
	if err != nil {
		return nil, err
	}

	token := init.JWTInstance.GenerateToken(user.Id.String(), role, true)
	return &token, nil
}
