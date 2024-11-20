package services

import "github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"

type IResourceCacheService interface {
	MakeKey(projectId, resourcePath string) string
	SetCache(resource *models.Resource) error
	GetCache(projectId, resourcePath string) (*models.ResourceCache, error)
	DeleteCache(projectId, resourcePath string)
}
