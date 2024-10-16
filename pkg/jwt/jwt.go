package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leandro-d-santos/no-code-api/config"
)

type Claims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(config.Env.JWTSecret)
}

func GenerateJWT(userId uint) (string, error) {
	oneDay := time.Hour * 24
	jwtDuration := oneDay * 7
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "no-code-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
