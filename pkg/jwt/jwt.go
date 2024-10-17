package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/pkg/cache"
)

type Claims struct {
	UserId uint   `json:"userId"`
	Stamp  string `json:"stamp"`
	jwt.RegisteredClaims
}

const (
	issuer   string = "https://node-code-api/api"
	audience string = "https://node-code-api"
	cacheKey string = "tokens"
)

func GetJWTSecret() []byte {
	return []byte(config.Env.JWTSecret)
}

func GenerateJWT(userId uint) (string, error) {
	oneDay := time.Hour * 24
	jwtDuration := oneDay * 7
	uuid := uuid.NewString()
	claims := &Claims{
		UserId: userId,
		Stamp:  uuid,
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

	if err := setStampInCache(userId, uuid); err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(userToken string) (*Claims, string) {
	jwtSecret := GetJWTSecret()
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}

	token, err := jwt.ParseWithClaims(userToken, &Claims{}, keyFunc)
	if err != nil {
		return nil, "Erro ao interpretar reivindicações"
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid || claims.Stamp == "" {
		return nil, "Token inválido"
	}

	cachedStamp, err := getStampFromCache(claims.UserId)
	if err != nil || cachedStamp == "" {
		return nil, "Erro ao consultar carimbo salvo"
	}

	if cachedStamp != claims.Stamp {
		return nil, "Carimbo fornecido é inválido"
	}

	return claims, ""
}

func getStampFromCache(userId uint) (string, error) {
	key := buildKey(userId)
	return cache.Get(key)
}

func setStampInCache(userId uint, stamp string) error {
	key := buildKey(userId)
	return cache.Set(key, stamp)
}

func buildKey(userId uint) string {
	return cacheKey + ":" + strconv.FormatUint(uint64(userId), 10)
}
