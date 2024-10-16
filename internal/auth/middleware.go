package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		claims, valid := internalJwt.ValidateToken(headerToken)
		if !valid {
			utils.ResJson(c, http.StatusUnauthorized, false, "Token inválido", nil)
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
