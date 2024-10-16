package users

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
)

func RegisterUsersRoutesV1(rg *gin.RouterGroup) {
	rg.POST("/login", HandleLogin)
	routes := rg.Group("/users")
	{
		routes.POST("/", HandleCreate)
		routes.Use(auth.AuthMiddleware())
		routes.GET("/", HandleFindAll)
		routes.GET("/:id", HandleFindById)
		routes.PUT("/:id", HandleUpdate)
		routes.DELETE("/:id", HandleDelete)
	}
}
