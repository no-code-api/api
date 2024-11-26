package external_endpoint

import (
	"fmt"
	"strings"

	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/services"
)

type ExternalEndpointService struct {
	resourceCacheService       services.IResourceCacheService
	resourceDynamicDataService services.IResourceDynamicDataService
}

type IExternalEndpointService interface {
	Get(request request) (interface{}, error)
	Post(request request) (interface{}, error)
}

func NewService() IExternalEndpointService {
	return ExternalEndpointService{
		resourceCacheService:       services.NewService(),
		resourceDynamicDataService: services.NewResourceDynamicDataService(),
	}
}

func (s ExternalEndpointService) Get(request request) (interface{}, error) {
	resourceCache, err := s.findProjectEndpoint(request.ProjectId, request.Path)
	if err != nil {
		return nil, err
	}
	return resourceCache, err
}

func (s ExternalEndpointService) Post(request request) (interface{}, error) {
	return nil, nil
}

func (s ExternalEndpointService) findProjectEndpoint(projectId, path string) (*models.ResourceCache, error) {
	segments := strings.Split(path, "/")
	segmentsKey := "/"

	for _, segment := range segments {
		segmentsKey = fmt.Sprintf("%s/%s", segment)
		cache, err := s.resourceCacheService.GetCache(projectId, segmentsKey)
		if err != nil {
			return nil, err
		} else if cache.Exists {
			return cache, nil
		}
	}
	return nil, nil
}
