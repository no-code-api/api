package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-code-api/api/config"
	"github.com/no-code-api/api/internal/router"
)

func Initialize() {
	r := gin.Default()
	router.RegisterRoutes(r)
	r.Run(config.Env.ServerPort)
}
