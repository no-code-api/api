package services

import "github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"

type IResourceDynamicDataService interface {
	CreateCollection(projectId string) error
	Find(filter *models.ResourceDynamicFilter) ([]interface{}, error)
	Add(addModel *models.AddResourceDynamic) error
	Update(updateModel *models.UpdateResourceDynamic) error
	Delete(filter *models.ResourceDynamicFilter) error
	DropCollection(projectId string) error
}
