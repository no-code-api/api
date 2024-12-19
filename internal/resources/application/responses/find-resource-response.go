package responses

import (
	"github.com/no-code-api/api/internal/resources/domain/models"
)

type FindResourceResponse struct {
	Id        string                 `json:"id"`
	Path      string                 `json:"path"`
	Endpoints []FindEndpointResponse `json:"endpoints"`
}

func (frr *FindResourceResponse) FromModel(resource *models.Resource) {
	frr.Id = resource.Id
	frr.Path = resource.Path
	frr.Endpoints = make([]FindEndpointResponse, len(resource.Endpoints))
	for i, endpoint := range resource.Endpoints {
		endpointResponse := &FindEndpointResponse{}
		endpointResponse.FromModel(endpoint)
		frr.Endpoints[i] = *endpointResponse
	}
}
