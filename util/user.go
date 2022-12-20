package util

import (
	customJwt "github.com/floatkasemtan/authentacle-service/init"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserInfo(c *gin.Context) (*string, *int8, *bool, error) {
	const BearerSchema = "Bearer "
	authKey := c.GetHeader("Authorization")
	tokenString := authKey[len(BearerSchema):]
	token, err := customJwt.JWTInstance.ValidateToken(tokenString)
	if !token.Valid {
		return nil, nil, nil, err
	}
	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"].(string)
	role := int8(claims["role"].(float64))
	verified := claims["verified"].(bool)

	return &id, &role, &verified, nil

}
