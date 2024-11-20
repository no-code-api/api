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
		endpoints[i] = createEndpoint.ToModel()
	}
	return endpoints
}
