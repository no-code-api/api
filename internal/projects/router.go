package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/auth"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	routes := rg.Group("/projects")
	routes.Use(auth.AuthMiddleware())
	routes.POST("/", handleCreate)
	routes.GET("/", handleFindByUser)
	routes.PUT("/", handleUpdate)
	routes.DELETE("/:id", handleDeleteByUser)
}
