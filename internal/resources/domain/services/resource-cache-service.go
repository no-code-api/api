package services

import (
	"encoding/json"
	"fmt"

	"github.com/no-code-api/api/internal/logger"
	"github.com/no-code-api/api/internal/resources/domain/models"
	"github.com/no-code-api/api/pkg/cache"
)

type cacheService struct {
	logger *logger.Logger
}

func NewService() IResourceCacheService {
	return &cacheService{
		logger: logger.NewLogger("ResourceCacheService"),
	}
}

func (s cacheService) SetCache(resource *models.Resource) error {
	resourceCache := &models.ResourceCache{
		Path:      resource.Path,
		Endpoints: make([]*models.EndpointCache, len(resource.Endpoints)),
	}
	for i, endpoint := range resource.Endpoints {
		endpointCache := &models.EndpointCache{
			Path:   endpoint.Path,
			Method: endpoint.Method,
		}
		resourceCache.Endpoints[i] = endpointCache
	}
	data, err := json.Marshal(resourceCache)
	if err != nil {
		return err
	}
	key := s.MakeKey(resource.ProjectId, resource.Path)
	if err := cache.Set(key, data); err != nil {
		s.logger.ErrorF("Error to set resource cache. %s", err.Error())
		return err
	}
	return nil
}

func (s cacheService) GetCache(projectId, resourcePath string) (*models.ResourceCache, error) {
	key := s.MakeKey(projectId, resourcePath)
	data, _ := cache.Get(key)
	if data == "" {
		return &models.ResourceCache{
			Exists:    false,
			Path:      "",
			Endpoints: nil,
		}, nil
	}
	resource := &models.ResourceCache{}
	if err := json.Unmarshal([]byte(data), resource); err != nil {
		return nil, fmt.Errorf("erro ao ler json do cache de recursos")
	}
	resource.Exists = true
	return resource, nil
}

func (s cacheService) DeleteCache(projectId, resourcePath string) {
	key := s.MakeKey(projectId, resourcePath)
	cache.Delete(key)
}

func (s cacheService) MakeKey(projectId, resourcePath string) string {
	return fmt.Sprintf("%s:%s", projectId, resourcePath)
}
