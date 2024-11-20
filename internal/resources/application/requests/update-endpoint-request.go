package requests

import "github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"

type UpdateEndpointRequest struct {
	Id     uint   `json:"id"`
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
