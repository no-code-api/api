package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

func RegisterRoutes(r *gin.Engine) {
	mainGroup := r.Group("/api")
	registerUsersRoutesV1(mainGroup)
}

func registerUsersRoutesV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	users.RegisterUsersRoutesV1(v1)
}
