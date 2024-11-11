package resources

type CreateEndpointRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (createEndpoint *CreateEndpointRequest) ToModel() *Endpoint {
	return &Endpoint{
		Path:   createEndpoint.Path,
		Method: createEndpoint.Method,
	}
}
