package router

import (
	"github.com/gin-gonic/gin"
	"github.com/no-code-api/api/internal/auth"
	"github.com/no-code-api/api/internal/handler"
	"github.com/no-code-api/api/internal/resources/application/handlers"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	endpointHandler := handlers.NewEndpointHandler()
	routes := rg.Group(endpointHandler.DefaultPath)
	routes.Use(auth.AuthMiddleware())
	routes.POST("/", handler.Wrapper(endpointHandler.HandleCreate))
	routes.GET("/", handler.Wrapper(endpointHandler.HandleFindAll))
	routes.PUT("/:resourceId", handler.Wrapper(endpointHandler.HandleUpdate))
	routes.DELETE("/:resourceId", handler.Wrapper(endpointHandler.HandleDelete))
}
