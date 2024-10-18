package database

import (
	"fmt"
	"time"

	"github.com/leandro-d-santos/no-code-api/config"
	internalLogger "github.com/leandro-d-santos/no-code-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connection *Connection
	logger     *internalLogger.Logger = internalLogger.NewLogger("Postgre")
)

type Connection struct {
	db *gorm.DB
}

func newConnection(gormDb *gorm.DB) *Connection {
	return &Connection{
		db: gormDb,
	}
}

func InitializePostgres() {
	host := config.Env.PostgreHost
	port := config.Env.PostgrePort
	user := config.Env.PostgreUserName
	password := config.Env.PostgrePassword
	dbName := config.Env.PostgreDbName
	sslMode := config.Env.PostgreSSLMode

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC", host, port, user, password, dbName, sslMode)
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.FatalF("Fail to connect on database: %v", err)
	}

	RunMigrations(gormDb)

	sqlDB, err := gormDb.DB()
	if err != nil {
		logger.FatalF("Fail to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	connection = newConnection(gormDb)
	logger.Debug("Database initialized")
}

func GetConnection() *Connection {
	if connection == nil {
		logger.Fatal("Database not initialized")
	}
	return connection
}

func (c *Connection) Save(data interface{}, exists bool) (ok bool) {
	var result *gorm.DB
	if !exists {
		result = c.db.Create(data)
	} else {
		result = c.db.Updates(data)
	}
	if result.Error != nil {
		logger.ErrorF("Error to save data: %v", result.Error.Error())
	}
	return result.Error == nil
}

func (c *Connection) Find(dest interface{}, conds ...interface{}) (ok bool) {
	if result := c.db.Find(dest, conds...); result.Error != nil {
		logger.ErrorF("Error to search data: filters: %v error: %v", conds, result.Error.Error())
		return false
	}
	return true
}

func (c *Connection) Delete(data interface{}, conds ...interface{}) (ok bool) {
	result := c.db.Delete(data, conds...)
	if result.Error != nil {
		logger.ErrorF("Error to delete data: filters: %v error: %v", conds, result.Error.Error())
	}
	return result.Error == nil
}
