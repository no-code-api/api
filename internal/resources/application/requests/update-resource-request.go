package requests

type UpdateResourceRequest struct {
	Id        string                   `json:"id"`
	Path      string                   `json:"path"`
	Endpoints []*UpdateEndpointRequest `json:"endpoints"`
	ProjectId string
}
