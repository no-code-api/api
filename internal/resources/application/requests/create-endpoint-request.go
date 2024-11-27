package requests

import "github.com/no-code-api/api/internal/resources/domain/models"

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
