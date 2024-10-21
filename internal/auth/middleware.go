package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/jwt"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
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
		service := jwt.NewJwtService()
		claims, err := service.ValidateToken(headerToken)
		if err != "" {
			msg := "Token inválido: " + err
			utils.ResJson(c, http.StatusUnauthorized, false, msg, nil)
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
