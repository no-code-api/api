package services

import (
	"fmt"
	"strings"

	"github.com/leandro-d-santos/no-code-api/internal/external-endpoint/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/core"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/services"
)

type externalEndpointService struct {
	resourceCacheService       services.IResourceCacheService
	resourceDynamicDataService services.IResourceDynamicDataService
}

func NewService() IExternalEndpointService {
	return externalEndpointService{
		resourceCacheService:       services.NewService(),
		resourceDynamicDataService: services.NewResourceDynamicDataService(),
	}
}

func (s externalEndpointService) Handle(request *requests.Request) (interface{}, error) {
	s.sanitizePaths(request)
	method := strings.ToUpper(request.Method)
	var data interface{}
	var err error
	switch method {
	case "GET":
		data, err = s.get(request)
	case "POST":
		data, err = s.post(request)
	default:
		return nil, fmt.Errorf("método '%s' não implementado", method)
	}
	return data, err
}

func (s externalEndpointService) get(request *requests.Request) (interface{}, error) {
	s.sanitizePaths(request)
	resourceCache, err := s.findCachedResource(request.ProjectId, request.Path)
	if err != nil {
		return nil, err
	}

	endpoints := s.getEndpointsByMethod(resourceCache.Endpoints, request.Method)
	fields := make([]models.ResourceDynamicFieldFilter, 0)
	requestPathSegments := strings.Split(request.Path, "/")
	for _, endpoint := range endpoints {
		endpointPath := core.SanitizeSuffixPath(resourceCache.Path + endpoint.Path)
		endpointPathSegments := strings.Split(endpointPath, "/")
		if len(endpointPathSegments) != len(requestPathSegments) {
			continue
		}
		for index, requestPathSegment := range requestPathSegments {
			endpointPathSegment := endpointPathSegments[index]
			hasId := strings.HasPrefix(endpointPathSegment, ":")
			if hasId {
				key := strings.TrimPrefix(endpointPathSegment, ":")
				val := requestPathSegment
				fields = append(fields, models.ResourceDynamicFieldFilter{Key: key, Value: val})
			} else if endpointPathSegment != requestPathSegment {
				break
			}
		}
	}
	filter := &models.ResourceDynamicFilter{
		ProjectId:    request.ProjectId,
		ResourcePath: resourceCache.Path,
		Fields:       fields,
	}
	return s.resourceDynamicDataService.Find(filter)
}

func (s externalEndpointService) post(request *requests.Request) (interface{}, error) {
	return nil, nil
}

func (s externalEndpointService) getEndpointsByMethod(cachedEndpoints []*models.EndpointCache, method string) []*models.EndpointCache {
	var endpoints []*models.EndpointCache
	for _, endpoint := range cachedEndpoints {
		if endpoint.Method == method {
			endpoints = append(endpoints, endpoint)
		}
	}
	return endpoints
}

func (s externalEndpointService) findCachedResource(projectId, path string) (*models.ResourceCache, error) {
	segments := strings.Split(path, "/")
	segmentsKey := ""
	fmt.Println(segments)
	for _, segment := range segments {
		if segmentsKey == "/" {
			segmentsKey = ""
		}
		segmentsKey = fmt.Sprintf("%s/%s", segmentsKey, segment)
		cache, err := s.resourceCacheService.GetCache(projectId, segmentsKey)
		if err != nil {
			return nil, err
		} else if cache.Exists {
			return cache, nil
		}
	}
	return nil, fmt.Errorf("endpoint não encontrado")
}

func (s externalEndpointService) sanitizePaths(request *requests.Request) {
	request.Path = core.SanitizeSuffixPath(request.Path)
}
