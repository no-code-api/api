package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/cache"
	internalJWT "github.com/leandro-d-santos/no-code-api/pkg/jwt"
)

type JwtService struct {
	jwtSettings *internalJWT.JwtSettings
	logger      *logger.Logger
}

type Claims struct {
	UserId uint   `json:"userId"`
	Stamp  string `json:"stamp"`
	jwt.RegisteredClaims
}

func NewJwtService() *JwtService {
	return &JwtService{
		jwtSettings: internalJWT.NewJwtSettings(),
		logger:      logger.NewLogger("JwtService"),
	}
}

func (s *JwtService) GenerateJWT(userId uint) (string, error) {
	oneDay := time.Hour * 24
	jwtDuration := oneDay * 7
	uuid := uuid.NewString()
	claims := &Claims{
		UserId: userId,
		Stamp:  uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.jwtSettings.GetIssuer(),
			Audience:  jwt.ClaimStrings{s.jwtSettings.GetAudience()},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSettings.GetJWTSecret())
	if err != nil {
		return "", err
	}

	if err := s.setStampInCache(userId, uuid); err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) ValidateToken(userToken string) (*Claims, string) {
	jwtSecret := s.jwtSettings.GetJWTSecret()
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

	cachedStamp := s.getStampFromCache(claims.UserId)
	if cachedStamp == "" || cachedStamp != claims.Stamp {
		return nil, "O Carimbo fornecido é inválido"
	}

	return claims, ""
}

func (s *JwtService) RemoveStamp(userId uint) {
	key := s.jwtSettings.BuildKey(userId)
	cache.Delete(key)
}

func (s *JwtService) setStampInCache(userId uint, stamp string) error {
	key := s.jwtSettings.BuildKey(userId)
	return cache.Set(key, stamp)
}

func (s *JwtService) getStampFromCache(userId uint) string {
	key := s.jwtSettings.BuildKey(userId)
	val, _ := cache.Get(key)
	return val
}
