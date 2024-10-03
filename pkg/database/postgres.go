package database

import (
	"fmt"
	"time"

	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializePostgres() {
	host := config.Env.PostgreHost
	port := config.Env.PostgrePort
	user := config.Env.PostgreUserName
	password := config.Env.PostgrePassword
	dbName := config.Env.PostgreDbName
	sslMode := config.Env.PostgreSSLMode
	logger := logger.NewLogger("Postgre")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC", host, port, user, password, dbName, sslMode)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.FatalF("Fail to connect on database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.FatalF("Fail to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func GetDb() *gorm.DB {
	logger := logger.NewLogger("Postgre")
	if db != nil {
		logger.Fatal("Database not initialized")
	}
	return db
}
