package main

import (
	"github.com/leandro-d-santos/no-code-api/cmd/api"
	"github.com/leandro-d-santos/no-code-api/config"
)

func main() {
	config.Initialize()
	api.Initialize()
}
