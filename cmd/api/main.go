package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/internal/router"
)

func Initialize() {
	r := gin.Default()
	router.RegisterRoutes(r)
	r.Run(config.Env.ServerPort)
}
