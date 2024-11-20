package resources

type UpdateEndpointRequest struct {
	Id     uint   `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (endpoint *UpdateEndpointRequest) ToModel() *Endpoint {
	return &Endpoint{
		Id:     endpoint.Id,
		Path:   endpoint.Path,
		Method: endpoint.Method,
	}
}
