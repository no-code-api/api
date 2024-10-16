package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
	internalJwt "github.com/leandro-d-santos/no-code-api/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.ResJson(c, http.StatusUnauthorized, false, "Cabeçalho de autenticação não encontrado", nil)
			c.Abort()
			return
		}
		headerToken := strings.Split(authHeader, "Bearer ")[1]
		if headerToken == "" {
			utils.ResJson(c, http.StatusUnauthorized, false, "Formato do token inválido", nil)
			c.Abort()
			return
		}

		token, err := getToken(headerToken)
		if err != nil {
			utils.ResJson(c, http.StatusUnauthorized, false, "Token inválido", nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*internalJwt.Claims)
		if !ok || !token.Valid {
			utils.ResJson(c, http.StatusUnauthorized, false, "Token inválido", nil)
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}

func getToken(headerToken string) (*jwt.Token, error) {
	jwtSecret := internalJwt.GetJWTSecret()
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}
	return jwt.ParseWithClaims(
		headerToken,
		&internalJwt.Claims{},
		keyFunc)
}
