package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
)

type Config struct {
	ServerPort string
}

var Env *Config

func Initialize() {
	logger := logger.NewLogger("Config")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	Env = &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

func GetLogger(prefix string) *logger.Logger {
	return logger.NewLogger(prefix)
}
