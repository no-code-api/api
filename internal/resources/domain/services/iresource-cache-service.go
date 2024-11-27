package services

import "github.com/no-code-api/no-code-api/internal/resources/domain/models"

type IResourceCacheService interface {
	MakeKey(projectId, resourcePath string) string
	SetCache(resource *models.Resource) error
	GetCache(projectId, resourcePath string) (*models.ResourceCache, error)
	DeleteCache(projectId, resourcePath string)
}
