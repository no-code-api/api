package services

import (
	"errors"
	"fmt"
	"reflect"
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
	case "PUT":
		data, err = s.put(request)
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
	endpoint := s.findCachedEndpoint(resourceCache.Path, request.Path, endpoints)
	if endpoint == nil {
		return nil, fmt.Errorf("endpoint '%s' para o método '%s' não encontrado", request.Method, request.Path)
	}
	endpointPath := core.SanitizeSuffixPath(resourceCache.Path + endpoint.Path)
	filter := &models.ResourceDynamicFilter{
		ProjectId:    request.ProjectId,
		ResourcePath: resourceCache.Path,
		Fields:       s.getFilterFields(endpointPath, request.Path),
	}
	return s.resourceDynamicDataService.Find(filter)
}

func (externalEndpointService) getFilterFields(endpointPath, requestPath string) []models.ResourceDynamicFieldFilter {
	fields := make([]models.ResourceDynamicFieldFilter, 0)
	endpointPathSegments := strings.Split(endpointPath, "/")
	requestPathSegments := strings.Split(requestPath, "/")
	for index, requestPathSegment := range requestPathSegments {
		endpointPathSegment := endpointPathSegments[index]
		hasId := strings.HasPrefix(endpointPathSegment, ":")
		if hasId {
			key := strings.TrimPrefix(endpointPathSegment, ":")
			val := requestPathSegment
			fields = append(fields, models.ResourceDynamicFieldFilter{Key: key, Value: val})
		}
	}
	return fields
}

func (s externalEndpointService) post(request *requests.Request) (interface{}, error) {
	s.sanitizePaths(request)
	resourceCache, err := s.findCachedResource(request.ProjectId, request.Path)
	if err != nil {
		return nil, err
	}

	endpoints := s.getEndpointsByMethod(resourceCache.Endpoints, request.Method)
	if endpoint := s.findCachedEndpoint(resourceCache.Path, request.Path, endpoints); endpoint == nil {
		return nil, fmt.Errorf("endpoint '%s' para o método '%s' não encontrado", request.Method, request.Path)
	}

	if err := s.validBody(request); err != nil {
		return nil, err
	}

	rows := s.transformBodyToRows(request)
	values := &models.AddResourceDynamic{
		ProjectId:    request.ProjectId,
		ResourcePath: resourceCache.Path,
		Rows:         rows,
	}
	if err := s.resourceDynamicDataService.Add(values); err != nil {
		return nil, fmt.Errorf("erro ao cadastrar valores. %s", err)
	}
	return nil, nil
}

func (s externalEndpointService) put(request *requests.Request) (interface{}, error) {
	s.sanitizePaths(request)
	resourceCache, err := s.findCachedResource(request.ProjectId, request.Path)
	if err != nil {
		return nil, err
	}

	endpoints := s.getEndpointsByMethod(resourceCache.Endpoints, request.Method)
	endpoint := s.findCachedEndpoint(resourceCache.Path, request.Path, endpoints)
	if endpoint == nil {
		return nil, fmt.Errorf("endpoint '%s' para o método '%s' não encontrado", request.Method, request.Path)
	}

	if err := s.validBody(request); err != nil {
		return nil, err
	}

	endpointPath := core.SanitizeSuffixPath(resourceCache.Path + endpoint.Path)
	values := &models.UpdateResourceDynamic{
		ProjectId:    request.ProjectId,
		ResourcePath: resourceCache.Path,
		Fields:       s.getFilterFields(endpointPath, request.Path),
		Data:         request.Body,
	}
	if err := s.resourceDynamicDataService.Update(values); err != nil {
		return nil, fmt.Errorf("erro ao cadastrar valores. %s", err)
	}
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

func (s externalEndpointService) findCachedEndpoint(resourceCachePath, requestPath string, endpoints []*models.EndpointCache) *models.EndpointCache {
	requestPathSegments := strings.Split(requestPath, "/")
	var endpointResult *models.EndpointCache = nil
	for _, endpoint := range endpoints {
		endpointResult = endpoint
		endpointPath := core.SanitizeSuffixPath(resourceCachePath + endpoint.Path)
		endpointPathSegments := strings.Split(endpointPath, "/")
		if len(endpointPathSegments) != len(requestPathSegments) {
			endpointResult = nil
			continue
		}
		fmt.Println("Endpoint")
		for index, requestPathSegment := range requestPathSegments {
			endpointPathSegment := endpointPathSegments[index]
			hasId := strings.HasPrefix(endpointPathSegment, ":")
			if !hasId && endpointPathSegment != requestPathSegment {
				endpointResult = nil
				break
			}
		}
		if endpointResult != nil {
			break
		}
	}
	return endpointResult
}

func (s externalEndpointService) validBody(request *requests.Request) error {
	canAddList := false
	reflectKind := reflect.TypeOf(request.Body).Kind()
	if !canAddList && (reflectKind == reflect.Slice || reflectKind == reflect.Array) {
		return errors.New("você não tem permissão para salvar uma lista de valores")
	}
	return nil
}

func (s externalEndpointService) transformBodyToRows(request *requests.Request) []interface{} {
	rows := make([]interface{}, 0)
	reflectKind := reflect.TypeOf(request.Body).Kind()
	if reflectKind == reflect.Slice || reflectKind == reflect.Array {
		rows = append(rows, request.Body.([]interface{})...)
	} else {
		rows = append(rows, request.Body)
	}
	return rows
}
