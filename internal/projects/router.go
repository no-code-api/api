package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	routes := rg.Group("/projects")
	routes.Use(auth.AuthMiddleware())
	routes.POST("/", handler.Wrapper(handleCreate))
	routes.GET("/", handler.Wrapper(handleFindByUser))
	routes.PUT("/", handler.Wrapper(handleUpdate))
	routes.DELETE("/:projectId", handler.Wrapper(handleDeleteByUser))
}
