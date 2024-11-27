package requests

import "github.com/no-code-api/no-code-api/internal/resources/domain/models"

type UpdateEndpointRequest struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (endpoint *UpdateEndpointRequest) ToModel() *models.Endpoint {
	return &models.Endpoint{
		Id:     endpoint.Id,
		Path:   endpoint.Path,
		Method: endpoint.Method,
	}
}
