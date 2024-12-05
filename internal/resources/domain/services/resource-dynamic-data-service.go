package services

import (
	"context"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/core"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
	"github.com/leandro-d-santos/no-code-api/pkg/mongodb"
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

func (s ResourceDynamicDataService) Find(filter *models.ResourceDynamicFilter) ([]interface{}, error) {
	collectionName := core.GetCollectionName(filter.ProjectId)
	mongFilter := s.buildFilter(filter)
	cur, err := s.mongoClient.Collection(collectionName).Find(context.Background(), mongFilter)
	if err != nil {
		fmt.Println("Erro: ", err)
		return nil, fmt.Errorf("erro ao consultar os dados do caminho '%s'", filter.ResourcePath)
	}

	var rows []struct{ Data interface{} }
	if err = cur.All(context.Background(), &rows); err != nil {
		return nil, err
	}

	var results []interface{} = make([]interface{}, len(rows))
	for i, row := range rows {
		results[i] = row.Data
	}
	return results, nil
}

func (s ResourceDynamicDataService) buildFilter(filter *models.ResourceDynamicFilter) bson.D {
	mongFilter := bson.D{{Key: "resourcePath", Value: filter.ResourcePath}}
	if filter.Fields != nil {
		for _, filter := range filter.Fields {
			key := fmt.Sprintf("data.%s", filter.Key)
			val := filter.Value
			mongFilter = append(mongFilter, bson.E{Key: key, Value: val})
		}
	}
	return mongFilter
}

func (s ResourceDynamicDataService) DropCollection(projectId string) error {
	collectionName := core.GetCollectionName(projectId)
	return s.mongoClient.Collection(collectionName).Drop(context.Background())
}
