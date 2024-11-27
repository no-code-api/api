package router

import (
	"github.com/gin-gonic/gin"
	"github.com/no-code-api/no-code-api/internal/auth"
	"github.com/no-code-api/no-code-api/internal/handler"
	"github.com/no-code-api/no-code-api/internal/users/application/handlers"
)

func RegisterUsersRoutesV1(rg *gin.RouterGroup) {
	userHandler := handlers.NewHandler()
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
