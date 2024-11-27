package responses

import "github.com/no-code-api/no-code-api/internal/resources/domain/models"

type FindEndpointResponse struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (ep *FindEndpointResponse) FromModel(endpoint *models.Endpoint) {
	ep.Id = endpoint.Id
	ep.Path = endpoint.Path
	ep.Method = endpoint.Method
}
