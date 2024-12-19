package services

import "github.com/no-code-api/api/internal/resources/domain/models"

type IResourceDynamicDataService interface {
	CreateCollection(projectId string) error
	DropCollection(projectId string) error
	UpdateResourcePath(projectId, resourcePath string) error
	Find(filter *models.ResourceDynamicFilter) ([]interface{}, error)
	Add(addModel *models.AddResourceDynamic) error
	Update(updateModel *models.UpdateResourceDynamic) error
	Delete(filter *models.ResourceDynamicFilter) error
}
