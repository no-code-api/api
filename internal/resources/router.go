package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	endpointHandler := NewEndpointHandler()
	routes := rg.Group(endpointHandler.DefaultPath)
	routes.Use(auth.AuthMiddleware())
	routes.POST("/", handler.Wrapper(endpointHandler.HandleCreate))
	routes.GET("/", handler.Wrapper(endpointHandler.HandleFindAll))
	routes.PUT("/:resourceId", handler.Wrapper(endpointHandler.HandleUpdate))
	routes.DELETE("/:resourceId", handler.Wrapper(endpointHandler.HandleDelete))
}
