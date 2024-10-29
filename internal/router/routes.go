package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/endpoints"
	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

func RegisterRoutes(r *gin.Engine) {
	mainGroup := r.Group("/api")
	registerRoutesV1(mainGroup)
}

func registerRoutesV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	users.RegisterUsersRoutesV1(v1)
	projects.RegisterRoutesV1(v1)
	endpoints.RegisterRoutesV1(v1)
}
