package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
	"github.com/leandro-d-santos/no-code-api/internal/projects/application/handlers"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	routes := rg.Group("/projects")
	routes.Use(auth.AuthMiddleware())
	projectHandler := handlers.NewHandler()
	routes.POST("/", handler.Wrapper(projectHandler.HandleCreate))
	routes.GET("/", handler.Wrapper(projectHandler.HandleFindByUser))
	routes.PUT("/:projectId", handler.Wrapper(projectHandler.HandleUpdate))
	routes.DELETE("/:projectId", handler.Wrapper(projectHandler.HandleDeleteByUser))
}
