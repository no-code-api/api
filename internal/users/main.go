package users

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

func RegisterUsersRoutesV1(rg *gin.RouterGroup) {
	rg.POST("/login", handler.Wrapper(HandleLogin))
	routes := rg.Group("/users")
	{
		routes.POST("/", handler.Wrapper(HandleCreate))
		routes.Use(auth.AuthMiddleware())
		routes.GET("/", handler.Wrapper(HandleFindAll))
		routes.GET("/:id", handler.Wrapper(HandleFindById))
		routes.PUT("/:id", handler.Wrapper(HandleUpdate))
		routes.DELETE("/:id", handler.Wrapper(HandleDelete))
	}
}
