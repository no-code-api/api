package main

import (
	"github.com/no-code-api/no-code-api/cmd/api"
	"github.com/no-code-api/no-code-api/config"
	"github.com/no-code-api/no-code-api/pkg/cache"
	"github.com/no-code-api/no-code-api/pkg/mongodb"
	"github.com/no-code-api/no-code-api/pkg/postgre"
)

func main() {
	config.Initialize()
	cache.InitializeRedis()
	postgre.InitializePostgres()
	mongodb.InitializeMongoDb()
	api.Initialize()
}
