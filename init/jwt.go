package jwt

import (
	"errors"
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTService interface {
	GenerateToken(id string, role int8, verified bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type CustomClaims struct {
	Id       string `json:"id"`
	Role     int8   `json:"role"`
	Verified bool   `json:"verified"`
	jwt.RegisteredClaims
}

type JWTServices struct {
	issuer string
	secret string
}

func NewJWTService() JWTService {
	return JWTServices{
		issuer: "Authentacle",
		secret: config.C.JWT_SECRET,
	}
}

func (service JWTServices) GenerateToken(id string, role int8, verified bool) string {
	claims := &CustomClaims{
		Id:       id,
		Role:     role,
		Verified: verified,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: service.issuer,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 72),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secret))
	if err != nil {
		panic(err)
	}
	return t
}

func (service JWTServices) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(service.secret), nil
	})
}

var JWTInstance JWTService

func init() {
	JWTInstance = NewJWTService()
}
