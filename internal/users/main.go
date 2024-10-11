package users

import "github.com/gin-gonic/gin"

func RegisterUsersRoutesV1(rg *gin.RouterGroup) {
	rg.POST("/login", HandleLogin)
	routes := rg.Group("/users")
	{
		routes.POST("/", HandleCreate)
		routes.GET("/", HandleFindAll)
		routes.GET("/:id", HandleFindById)
		routes.PUT("/:id", HandleUpdate)
		routes.DELETE("/:id", HandleDelete)
	}
}
