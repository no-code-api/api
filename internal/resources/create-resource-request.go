package resources

type CreateResourceRequest struct {
	Path      string                   `json:"path"`
	Endpoints []*CreateEndpointRequest `json:"endpoints"`
	ProjectId string
}

func (createResource *CreateResourceRequest) ToModel() *Resource {
	resource := &Resource{
		Path:      createResource.Path,
		ProjectId: createResource.ProjectId,
	}
	resource.Endpoints = createResource.EndpointsToModel(resource)
	return resource
}

func (createResource *CreateResourceRequest) EndpointsToModel(resource *Resource) []*Endpoint {
	endpoints := make([]*Endpoint, len(createResource.Endpoints))
	for i, createEndpoint := range createResource.Endpoints {
		endpoint := createEndpoint.ToModel()
		endpoint.Resource = resource
		endpoints[i] = endpoint
	}
	return endpoints
}
