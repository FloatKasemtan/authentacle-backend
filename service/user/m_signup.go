package user

import (
	jwt "github.com/floatkasemtan/authentacle-service/init"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func (s userService) SignUp(username string, email string, password string, userAgent string) (*string, *string, *string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return nil, nil, nil, err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Authentacle",
		AccountName: email,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	encryptedSecret, err := util.Encrypt(key.Secret())
	if err != nil {
		return nil, nil, nil, err
	}

	userId, err := s.userRepository.CreateUser(username, email, string(hashedPassword), encryptedSecret)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create JWT Token
	token := jwt.JWTInstance.GenerateToken(userId, 0, false, userAgent)
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
