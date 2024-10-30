package endpoints

type CreateEndpointRequest struct {
	Path      string `json:"path"`
	Method    string `json:"method"`
	ProjectId string
}

func (ev *CreateEndpointRequest) ToModel() *Endpoint {
	return &Endpoint{
		Path:      ev.Path,
		Method:    ev.Method,
		ProjectId: ev.ProjectId,
	}
}
