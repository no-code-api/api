package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/config"
)

func Initialize() {
	r := gin.Default()
	r.Run(config.Env.ServerPort)
}
