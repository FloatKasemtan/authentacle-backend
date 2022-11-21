package init

import (
	"errors"
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTService interface {
	GenerateToken(email string, role int8) string
	ValidateToken(token string) (*jwt.Token, error)
}

type CustomClaims struct {
	Email string `json:"email"`
	Role  int8   `json:"role"`
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

func (service JWTServices) GenerateToken(email string, role int8) string {
	claims := &CustomClaims{
		Email: email,
		Role:  role,
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
