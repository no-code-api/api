package endpoints

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
	routes.PUT("/:endpointId", handler.Wrapper(endpointHandler.HandleUpdate))
	routes.DELETE("/:endpointId", handler.Wrapper(endpointHandler.HandleDelete))
}
