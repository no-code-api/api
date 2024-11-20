package validations

import (
	"fmt"
	"strings"

	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/constants"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
)

func CreateResourceIsValid(resource *models.Resource) error {
	if err := validatePathLength("Caminho", resource.Path); err != nil {
		return err
	}
	if err := validateEmptyString("Código projeto", resource.ProjectId); err != nil {
		return err
	}
	if err := ValidateEndpoints(resource.Endpoints); err != nil {
		return err
	}
	return nil
}

func UpdateResourceIsValid(resource *models.Resource) error {
	if err := validateEmptyString("Código", resource.Id); err != nil {
		return err
	}
	if err := validatePathLength("Caminho", resource.Path); err != nil {
		return err
	}
	if err := validateEmptyString("Código projeto", resource.ProjectId); err != nil {
		return err
	}
	if err := ValidateEndpoints(resource.Endpoints); err != nil {
		return err
	}
	return nil
}

func validateEmptyString(propertyName, value string) error {
	if value == "" {
		return fmt.Errorf("'%s' não pode ser vazio", propertyName)
	}
	return nil
}

func validatePathLength(propertyName, path string) error {
	pathLen := len(path)
	if pathLen > 50 {
		return fmt.Errorf("'%s' dever ter 50 ou menos caracteres", propertyName)
	}
	return nil
}

func validateMethod(propertyName, method string) error {
	var allowedMethods = []string{constants.GET, constants.POST, constants.PUT, constants.DELETE}
	for _, allowedMethod := range allowedMethods {
		if allowedMethod == method {
			return nil
		}
	}
	return fmt.Errorf("'%s' dever estar entre: GET, POST, PUT, DELETE", propertyName)
}

func ValidateEndpoints(endpoints []*models.Endpoint) error {
	allPathsByMethod := make(map[string][]*models.Endpoint)
	for i, endpoint := range endpoints {
		methodProperty := fmt.Sprintf("Endpoint.[%d].Método", i)
		pathProperty := fmt.Sprintf("Endpoint.[%d].Caminho", i)
		if err := validateMethod(methodProperty, endpoint.Method); err != nil {
			return err
		}
		if err := validatePathLength(pathProperty, endpoint.Path); err != nil {
			return err
		}
		pathsByMethod, ok := allPathsByMethod[endpoint.Method]
		if !ok {
			pathsByMethod = make([]*models.Endpoint, 0)
			allPathsByMethod[endpoint.Method] = pathsByMethod
		}
		if err := validatePathSegment(endpoint, pathsByMethod); err != nil {
			return err
		}
		pathsByMethod = append(pathsByMethod, endpoint)
		allPathsByMethod[endpoint.Method] = pathsByMethod
	}
	return nil
}

func validatePathSegment(endpoint *models.Endpoint, pathsByMethod []*models.Endpoint) error {
	endpointSegments := strings.Split(endpoint.Path, "/")
	for _, path := range pathsByMethod {

		if path.Path == endpoint.Path && path.Id != endpoint.Id {
			return fmt.Errorf("endpoint já cadastrado: ('%s' - '%s')", endpoint.Method, endpoint.Path)
		}

		segments := strings.Split(path.Path, "/")
		if PathsConflict(endpointSegments, segments) {
			return fmt.Errorf("conflito detectado entre os endpoints: ('%s' - '%s') e ('%s' - '%s')", endpoint.Method, endpoint.Path, path.Method, path.Path)
		}
	}
	return nil
}

func PathsConflict(endpointSegments []string, pathSegments []string) bool {
	for i := 1; i < len(endpointSegments); i++ {
		originalEndpointSegment := endpointSegments[i]
		if i > (len(pathSegments) - 1) {
			return false
		}
		originalPathSegment := pathSegments[i]
		pathIsSame := originalEndpointSegment == originalPathSegment
		segmentContainsTwoDots := strings.HasPrefix(originalEndpointSegment, ":")
		pathSegmentContainsTwoDots := strings.HasPrefix(originalPathSegment, ":")
		if pathIsSame || (!segmentContainsTwoDots && !pathSegmentContainsTwoDots) {
			continue
		}
		if segmentContainsTwoDots && pathSegmentContainsTwoDots {
			return true
		}
		ok := validatePathIdParams(originalEndpointSegment, originalPathSegment)
		if !ok {
			return true
		}
	}
	return false
}

func validatePathIdParams(originalEndpointSegment string, originalPathSegment string) bool {
	segmentContainsTwoDots := strings.HasPrefix(originalEndpointSegment, ":")
	pathSegmentContainsTwoDots := strings.HasPrefix(originalPathSegment, ":")
	if segmentContainsTwoDots {
		originalEndpointSegment = ""
	}
	if pathSegmentContainsTwoDots {
		originalPathSegment = ""
	}
	return originalEndpointSegment == originalPathSegment
}
