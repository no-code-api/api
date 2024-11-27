package mongodb

import (
	"context"
	"fmt"

	"github.com/no-code-api/no-code-api/config"
	internalLogger "github.com/no-code-api/no-code-api/internal/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	connection *mongo.Database
	logger     *internalLogger.Logger = internalLogger.NewLogger("Mongodb")
)

func InitializeMongoDb() {
	userName := config.Env.MongoDbUserName
	password := config.Env.MongoDbPassword
	host := config.Env.MongoDbHost
	port := config.Env.MongoDbPort
	databaseName := config.Env.MongoDbDbName
	dns := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", userName, password, host, port)
	fmt.Println(dns)
	clientOptions := options.Client().ApplyURI(dns)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logger.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		logger.Fatal(err)
	}

	connection = client.Database(databaseName)
	fmt.Println("MongoDB initialized")
}

func GetConnection() *mongo.Database {
	if connection == nil {
		logger.Fatal("MongoDB not initialized")
	}
	return connection
}
