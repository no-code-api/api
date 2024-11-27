package requests

import "github.com/no-code-api/no-code-api/internal/resources/domain/models"

type CreateResourceRequest struct {
	Path      string                   `json:"path"`
	Endpoints []*CreateEndpointRequest `json:"endpoints"`
	ProjectId string
}

func (createResource *CreateResourceRequest) ToModel() *models.Resource {
	resource := &models.Resource{
		Path:      createResource.Path,
		ProjectId: createResource.ProjectId,
	}
	resource.Endpoints = createResource.EndpointsToModel(resource)
	return resource
}

func (createResource *CreateResourceRequest) EndpointsToModel(resource *models.Resource) []*models.Endpoint {
	endpoints := make([]*models.Endpoint, len(createResource.Endpoints))
	for i, createEndpoint := range createResource.Endpoints {
		endpoints[i] = createEndpoint.ToModel()
	}
	return endpoints
}
