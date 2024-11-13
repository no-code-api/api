package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/leandro-d-santos/no-code-api/config"
)

func InitializePostgres() {
	host := config.Env.PostgreHost
	port := config.Env.PostgrePort
	user := config.Env.PostgreUserName
	password := config.Env.PostgrePassword
	dbName := config.Env.PostgreDbName
	sslMode := config.Env.PostgreSSLMode

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC", host, port, user, password, dbName, sslMode)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		logger.FatalF("Fail to connect on database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := conn.Ping(ctx); err != nil {
		logger.FatalF("Failed to ping database: %v", err)
	}

	connection = newConnection(conn)
	RunMigrations(connection)
	logger.Debug("Database initialized")
}

func GetConnection() *Connection {
	if connection == nil {
		logger.Fatal("Database not initialized")
	}
	return connection
}
