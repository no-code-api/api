package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/no-code-api/api/internal/logger"
)

func LoadEnv() *Config {
	logger := logger.NewLogger("Config")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	return &Config{
		ServerPort:      os.Getenv("SERVER_PORT"),
		PostgreHost:     os.Getenv("POSTGRE_HOST"),
		PostgrePort:     os.Getenv("POSTGRE_PORT"),
		PostgreUserName: os.Getenv("POSTGRE_USER_NAME"),
		PostgrePassword: os.Getenv("POSTGRE_PASSWORD"),
		PostgreDbName:   os.Getenv("POSTGRE_DB_NAME"),
		PostgreSSLMode:  os.Getenv("POSTGRE_SSL_MODE"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		RedisHost:       os.Getenv("REDIS_HOST"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
		RedisDb:         os.Getenv("REDIS_DB"),
		MongoDbUserName: os.Getenv("MONGO_DB_USER_NAME"),
		MongoDbPassword: os.Getenv("MONGO_DB_PASSWORD"),
		MongoDbHost:     os.Getenv("MONGO_DB_HOST"),
		MongoDbPort:     os.Getenv("MONGO_DB_PORT"),
		MongoDbDbName:   os.Getenv("MONGO_DB_DB_NAME"),
		InternalDomain:  os.Getenv("INTERNAL_DOMAIN"),
	}
}
