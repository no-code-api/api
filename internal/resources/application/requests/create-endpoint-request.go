package requests

import "github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"

type CreateEndpointRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (createEndpoint *CreateEndpointRequest) ToModel() *models.Endpoint {
	return &models.Endpoint{
		Path:   createEndpoint.Path,
		Method: createEndpoint.Method,
	}
}
