package main

import (
	"github.com/leandro-d-santos/no-code-api/cmd/api"
	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

func main() {
	config.Initialize()
	database.InitializePostgres()
	api.Initialize()
}
