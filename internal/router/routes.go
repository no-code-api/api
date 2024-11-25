package router

import (
	"github.com/gin-gonic/gin"
	projectsRouter "github.com/leandro-d-santos/no-code-api/internal/projects/application/router"
	resourceRouter "github.com/leandro-d-santos/no-code-api/internal/resources/application/router"
	usersRouter "github.com/leandro-d-santos/no-code-api/internal/users/application/router"
)

func RegisterRoutes(r *gin.Engine) {
	mainGroup := r.Group("/api")
	registerRoutesV1(mainGroup)
}

func registerRoutesV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	usersRouter.RegisterUsersRoutesV1(v1)
	projectsRouter.RegisterRoutesV1(v1)
	resourceRouter.RegisterRoutesV1(v1)
}
