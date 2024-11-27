package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/no-code-api/no-code-api/internal/handler"
	"github.com/no-code-api/no-code-api/internal/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := handler.NewBaseHandler(c)
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			handler.Json(http.StatusUnauthorized, false, "Cabeçalho de autenticação não encontrado", nil)
			c.Abort()
			return
		}
		headerToken := strings.Split(authHeader, "Bearer ")[1]
		if headerToken == "" {
			handler.Json(http.StatusUnauthorized, false, "Formato do token inválido", nil)
			c.Abort()
			return
		}
		service := jwt.NewJwtService()
		claims, err := service.ValidateToken(headerToken)
		if err != "" {
			msg := "Token inválido: " + err
			handler.Json(http.StatusUnauthorized, false, msg, nil)
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
