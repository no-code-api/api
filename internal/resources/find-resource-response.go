package resources

type FindResourceResponse struct {
	Id        string                 `json:"id"`
	Path      string                 `json:"path"`
	Endpoints []FindEndpointResponse `json:"endpoints"`
}

func (frr *FindResourceResponse) FromModel(resource *Resource) {
	frr.Id = resource.Id
	frr.Path = resource.Path
	frr.Endpoints = make([]FindEndpointResponse, len(resource.Endpoints))
	for i, endpoint := range resource.Endpoints {
		endpointResponse := &FindEndpointResponse{}
		endpointResponse.FromModel(endpoint)
		frr.Endpoints[i] = *endpointResponse
	}
}
