package resources

type FindEndpointResponse struct {
	Id     uint   `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (ep *FindEndpointResponse) FromModel(endpoint *Endpoint) {
	ep.Id = endpoint.Id
	ep.Path = endpoint.Path
	ep.Method = endpoint.Method
}
