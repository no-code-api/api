package main

import (
	"github.com/no-code-api/api/cmd/api"
	"github.com/no-code-api/api/config"
	"github.com/no-code-api/api/pkg/cache"
	"github.com/no-code-api/api/pkg/mongodb"
	"github.com/no-code-api/api/pkg/postgre"
)

func main() {
	config.Initialize()
	cache.InitializeRedis()
	postgre.InitializePostgres()
	mongodb.InitializeMongoDb()
	api.Initialize()
}
