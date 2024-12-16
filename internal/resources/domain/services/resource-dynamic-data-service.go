package services

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/core"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
	"github.com/leandro-d-santos/no-code-api/pkg/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ResourceDynamicDataService struct {
	mongoClient *mongo.Database
	logger      *logger.Logger
}

type MongoRow struct {
	ResourcePath string `bson:"resourcePath"`
	Data         interface{}
}

var containsNumber = regexp.MustCompile(`\d`)

func NewResourceDynamicDataService() IResourceDynamicDataService {
	return ResourceDynamicDataService{
		mongoClient: mongodb.GetConnection(),
		logger:      logger.NewLogger("ResourceDynamicDataService"),
	}
}

func (s ResourceDynamicDataService) CreateCollection(projectId string) error {
	collectionName := core.GetCollectionName(projectId)
	if err := s.mongoClient.CreateCollection(context.Background(), collectionName); err != nil {
		return err
	}

	keys := bson.D{{Key: "resourcePath", Value: 1}}
	model := mongo.IndexModel{Keys: keys}
	if _, err := s.mongoClient.Collection(collectionName).Indexes().CreateOne(context.Background(), model); err != nil {
		return err
	}

	return nil
}

func (s ResourceDynamicDataService) Find(filter *models.ResourceDynamicFilter) ([]interface{}, error) {
	collectionName := core.GetCollectionName(filter.ProjectId)
	mongoFilter, err := s.buildFilter(filter.ResourcePath, filter.Fields)
	if err != nil {
		return nil, err
	}
	cur, err := s.mongoClient.Collection(collectionName).Find(context.Background(), mongoFilter)
	if err != nil {
		s.logger.DebugF("Error to search: MF(%v) Filter(%v) err(%s)", mongoFilter, filter, err)
		return nil, fmt.Errorf("erro ao consultar os dados do caminho '%s'", filter.ResourcePath)
	}

	var rows []struct{ Data interface{} }
	if err = cur.All(context.TODO(), &rows); err != nil {
		return nil, err
	}

	var results []interface{} = make([]interface{}, len(rows))
	for i, row := range rows {
		results[i] = row.Data
	}
	return results, nil
}

func (s ResourceDynamicDataService) buildFilter(resourcePath string, fields []models.ResourceDynamicFieldFilter) (bson.M, error) {
	mongoFilter := bson.M{"resourcePath": resourcePath}
	if len(fields) > 0 {
		andMapFilter := make(map[string][]bson.M)
		andFilter := make([]bson.M, 0)
		for _, filter := range fields {
			if err := s.addFilterValuesByField(andMapFilter, filter); err != nil {
				return nil, err
			}
		}
		for _, value := range andMapFilter {
			andFilter = append(andFilter, value...)
		}
		mongoFilter["$and"] = andFilter
	}
	return mongoFilter, nil
}

func (s ResourceDynamicDataService) addFilterValuesByField(andFilter map[string][]bson.M, filter models.ResourceDynamicFieldFilter) error {
	key := fmt.Sprintf("data.%s", filter.Key)
	val := filter.Value
	values, ok := andFilter[key]
	if !ok {
		values = make([]bson.M, 0)
	}
	if containsNumber.MatchString(val) {
		orFilterValues := make([]bson.M, 0)
		orFilterValues = append(orFilterValues, bson.M{key: val})
		var err error
		var numVal any = 0
		numVal, err = strconv.Atoi(filter.Value)
		if err != nil {
			numVal, err = strconv.ParseFloat(filter.Value, 64)
			if err != nil {
				return fmt.Errorf("erro ao converter parâmetro '%s'", filter.Key)
			}
		}
		orFilterValues = append(orFilterValues, bson.M{key: numVal})
		values = append(values, bson.M{"$or": orFilterValues})
	} else {
		values = append(values, bson.M{key: val})
	}
	andFilter[key] = values
	return nil
}

func (s ResourceDynamicDataService) Add(addModel *models.AddResourceDynamic) error {
	collectionName := core.GetCollectionName(addModel.ProjectId)
	rows := make([]*MongoRow, len(addModel.Rows))
	for i, addRow := range addModel.Rows {
		row := &MongoRow{
			ResourcePath: addModel.ResourcePath,
			Data:         addRow,
		}
		rows[i] = row
	}
	_, err := s.mongoClient.Collection(collectionName).InsertMany(context.TODO(), rows)
	if err != nil {
		s.logger.DebugF("Error to insert: Rows(%v) err(%s)", rows, err)
	}
	return err
}

func (s ResourceDynamicDataService) Update(updateModel *models.UpdateResourceDynamic) error {
	collectionName := core.GetCollectionName(updateModel.ProjectId)
	mongoFilter, err := s.buildFilter(updateModel.ResourcePath, updateModel.Fields)
	if err != nil {
		return err
	}

	row := s.buildRowToUpdate(updateModel)
	collection := s.mongoClient.Collection(collectionName)
	if _, err := collection.UpdateMany(context.TODO(), mongoFilter, bson.M{"$set": row}); err != nil {
		s.logger.DebugF("Error to update: MF(%v) Row(%v) err(%s)", mongoFilter, row, err)
		return err
	}
	return nil
}

func (s ResourceDynamicDataService) Delete(filter *models.ResourceDynamicFilter) error {
	collectionName := core.GetCollectionName(filter.ProjectId)
	mongoFilter, err := s.buildFilter(filter.ResourcePath, filter.Fields)
	if err != nil {
		return err
	}
	collection := s.mongoClient.Collection(collectionName)
	if _, err := collection.DeleteMany(context.TODO(), mongoFilter); err != nil {
		s.logger.DebugF("Error to delete: MF(%v) err(%s)", mongoFilter, err)
		return err
	}
	return nil
}

func (s ResourceDynamicDataService) UpdateResourcePath(projectId, resourcePath string) error {
	collectionName := core.GetCollectionName(projectId)
	collection := s.mongoClient.Collection(collectionName)
	mongoFilter := bson.M{}
	row := bson.M{
		"$set": bson.M{
			"resourcePath": resourcePath,
		},
	}
	if _, err := collection.UpdateMany(context.TODO(), mongoFilter, row); err != nil {
		s.logger.DebugF("Error to update resourcePath: ResourcePath(%s) ProjecrId(%s) err(%s)", resourcePath, projectId, err)
		return err
	}
	return nil
}

func (ResourceDynamicDataService) buildRowToUpdate(updateModel *models.UpdateResourceDynamic) bson.M {
	row := bson.M{}
	hasStruct := false
	if hasStruct {
		// não está completo, mas seria algo parecido a isso.
		// switch updateModel.Data.(type) {
		// case map[string]interface{}:
		// 	for key, a := range updateModel.Data.(map[string]interface{}) {
		// 		prop := fmt.Sprintf("data.%s", key)
		// 		row[prop] = a
		// 	}
		// default:
		// 	row["data"] = updateModel.Data
		// }
	} else {
		row["data"] = updateModel.Data
	}
	return row
}

func (s ResourceDynamicDataService) DropCollection(projectId string) error {
	collectionName := core.GetCollectionName(projectId)
	return s.mongoClient.Collection(collectionName).Drop(context.Background())
}
