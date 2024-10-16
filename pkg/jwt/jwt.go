package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
)

type Claims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

var (
	issuer   string = "https://node-code-api/api"
	audience string = "https://node-code-api"
)

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
			Issuer:    issuer,
			Audience:  jwt.ClaimStrings{audience},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(userToken string) (*Claims, bool) {
	jwtSecret := GetJWTSecret()
	logger := logger.NewLogger("ValidateToken")
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}

	token, err := jwt.ParseWithClaims(userToken, &Claims{}, keyFunc)
	if err != nil {
		logger.InfoF("Error to parse claims: %v", err.Error())
		return nil, false
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, false
	}

	return claims, true
}
