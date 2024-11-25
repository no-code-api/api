package repositories

import "github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"

type IRepository interface {
	CreateResource(resource *models.Resource) bool
	FindById(id string) (*models.Resource, bool)
	UpdateResource(resource *models.Resource) bool
	CheckResourcePathAvailableByProject(projectId string, path string) (bool, bool)
	CheckResourcePathAvailableByResourceId(resourceId string, path string) (bool, bool)
	FindAllResource(projectId string) ([]*models.Resource, bool)
	DeleteById(id string) bool
}
