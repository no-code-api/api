package main

import (
	"github.com/leandro-d-santos/no-code-api/cmd/api"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/pkg/cache"
	"github.com/leandro-d-santos/no-code-api/pkg/mongodb"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

func main() {
	config.Initialize()
	cache.InitializeRedis()
	postgre.InitializePostgres()
	mongodb.InitializeMongoDb()
	api.Initialize()
}
