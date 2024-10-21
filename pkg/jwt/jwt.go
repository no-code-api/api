package jwt

import (
	"strconv"

	"github.com/leandro-d-santos/no-code-api/config"
)

type JwtSettings struct {
	issuer    string
	audience  string
	cacheKey  string
	jwtSecret []byte
}

const (
	issuer   string = "https://node-code-api/api"
	audience string = "https://node-code-api"
	cacheKey string = "tokens"
)

func NewJwtSettings() *JwtSettings {
	return &JwtSettings{
		issuer:    issuer,
		audience:  audience,
		cacheKey:  cacheKey,
		jwtSecret: []byte(config.Env.JWTSecret),
	}
}

func (s *JwtSettings) GetJWTSecret() []byte {
	return s.jwtSecret
}

func (s *JwtSettings) GetCacheKey() string {
	return s.cacheKey
}

func (s *JwtSettings) GetIssuer() string {
	return s.issuer
}

func (s *JwtSettings) GetAudience() string {
	return s.audience
}

func (s *JwtSettings) BuildKey(id uint) string {
	return s.cacheKey + ":" + strconv.FormatUint(uint64(id), 10)
}
