package services

import (
	"context"

	"github.com/no-code-api/api/internal/resources/domain/core"
	"github.com/no-code-api/api/pkg/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ResourceDynamicDataService struct {
	mongoClient *mongo.Database
}

func NewResourceDynamicDataService() IResourceDynamicDataService {
	return ResourceDynamicDataService{
		mongoClient: mongodb.GetConnection(),
	}
}

func (s ResourceDynamicDataService) CreateCollection(projectId string) error {
	collectionName := core.GetCollectionName(projectId)
	if err := s.mongoClient.CreateCollection(context.Background(), collectionName); err != nil {
		return err
	}

	keys := bson.D{{"resourcePath", 1}}
	model := mongo.IndexModel{Keys: keys}
	if _, err := s.mongoClient.Collection(collectionName).Indexes().CreateOne(context.Background(), model); err != nil {
		return err
	}

	return nil
}

func (s ResourceDynamicDataService) DropCollection(projectId string) error {
	collectionName := core.GetCollectionName(projectId)
	return s.mongoClient.Collection(collectionName).Drop(context.Background())
}
