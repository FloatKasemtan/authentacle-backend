package user

import (
	jwt "github.com/floatkasemtan/authentacle-service/init"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

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
	token := jwt.JWTInstance.GenerateToken(user.ID.Hex(), 0, user.IsVerify)
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
