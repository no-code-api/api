package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/external-endpoint/application/handlers"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
	projectsRouter "github.com/leandro-d-santos/no-code-api/internal/projects/application/router"
	resourceRouter "github.com/leandro-d-santos/no-code-api/internal/resources/application/router"
	usersRouter "github.com/leandro-d-santos/no-code-api/internal/users/application/router"
)

func RegisterRoutes(mainServer *gin.Engine) {
	internalServer := gin.New()
	mainGroup := internalServer.Group("/api")
	registerRoutesV1(mainGroup)
	externalEndpointHandler := handlers.NewHandler()
	mainServer.Any("/*domain", func(ctx *gin.Context) {
		host := ctx.Request.Host
		if !strings.Contains(host, externalEndpointHandler.InternalDomain) {
			internalServer.HandleContext(ctx)
			return
		}
		externalEndpointHandler.Handle(handler.NewBaseHandler(ctx))
	})
}

func registerRoutesV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	usersRouter.RegisterUsersRoutesV1(v1)
	projectsRouter.RegisterRoutesV1(v1)
	resourceRouter.RegisterRoutesV1(v1)
}
