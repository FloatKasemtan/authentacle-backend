package util

import (
	"github.com/floatkasemtan/authentacle-service/init"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserInfo(c *gin.Context) (*string, *int8, *bool, error) {
	const BearerSchema = "Bearer "
	authKey := c.GetHeader("Authorization")
	tokenString := authKey[len(BearerSchema):]
	token, err := init.JWTInstance.ValidateToken(tokenString)
	if !token.Valid {
		return nil, nil, nil, err
	}
	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"].(string)
	level := claims["role"].(int8)
	verified := claims["verified"].(bool)

	return &id, &level, &verified, nil

}
