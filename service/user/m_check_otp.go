package user

import (
	"errors"
	jwt "github.com/floatkasemtan/authentacle-service/init"
	"github.com/floatkasemtan/authentacle-service/util"
	"github.com/pquerna/otp/totp"
)

func (s userService) CheckOTP(id string, role int8, otp string) (*string, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	decryptSecret, err := util.Decrypt(user.Secret)
	if err != nil {
		return nil, err
	}

	if !totp.Validate(otp, decryptSecret) {
		return nil, errors.New("Invalid OTP")
	}

	token := jwt.JWTInstance.GenerateToken(user.ID.Hex(), role, true)
	return &token, nil
}
