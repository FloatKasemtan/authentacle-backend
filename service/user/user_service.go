package user

import (
	"errors"
	"fmt"
	jwt "github.com/floatkasemtan/authentacle-service/init"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) userService {
	return userService{userRepository: userRepository}
}

func (s userService) SignUp(username string, email string, password string) (*string, *string, *string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return nil, nil, nil, err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Authentacle",
		AccountName: email,
	})

	userId, err := s.userRepository.SignUp(username, email, string(hashedPassword), key.Secret())
	if err != nil {
		return nil, nil, nil, err
	}

	// Create JWT Token
	token := jwt.JWTInstance.GenerateToken(userId, 0, true)
	if err != nil {
		return nil, nil, nil, err
	}

	if err != nil {
		return nil, nil, nil, err
	}

	secret := key.Secret()
	url := key.URL()
	return &token, &url, &secret, nil
}

func (s userService) SignIn(username string, password string) (*string, *bool, *string, *string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Test revert check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, nil, nil, err
	}

	// Create JWT Token
	token := jwt.JWTInstance.GenerateToken(user.Id.Hex(), 0, user.IsVerify)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if user.IsVerify {
		return &token, &user.IsVerify, nil, nil, nil
	} else {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Authentacle",
			AccountName: user.Email,
			Secret:      []byte(user.Secret),
		})
		if err != nil {
			return nil, nil, nil, nil, err
		}

		url := key.URL()
		return &token, &user.IsVerify, &user.Secret, &url, nil
	}
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

func (s userService) CheckOTP(id string, role int8, otp string) (*string, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.Secret)
	if !totp.Validate(otp, user.Secret) {
		return nil, errors.New("Invalid OTP")
	}

	token := jwt.JWTInstance.GenerateToken(user.Id.Hex(), role, true)
	return &token, nil
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

	token := jwt.JWTInstance.GenerateToken(user.Id.String(), role, true)
	return &token, nil
}
