package users

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

func RegisterUsersRoutesV1(rg *gin.RouterGroup) {
	userHandler := NewHandler()
	rg.POST("/login", handler.Wrapper(userHandler.HandleLogin))
	routes := rg.Group("/users")
	{
		routes.POST("/", handler.Wrapper(userHandler.HandleCreate))
		routes.Use(auth.AuthMiddleware())
		routes.GET("/", handler.Wrapper(userHandler.HandleFindAll))
		routes.GET("/:id", handler.Wrapper(userHandler.HandleFindById))
		routes.PUT("/:id", handler.Wrapper(userHandler.HandleUpdate))
		routes.DELETE("/:id", handler.Wrapper(userHandler.HandleDelete))
	}
}
