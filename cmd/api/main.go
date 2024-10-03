package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

func Initialize() {
	database.InitializePostgres()
	r := gin.Default()
	r.Run(config.Env.ServerPort)
}
