package endpoints

type UpdateEndpointRequest struct {
	Id        uint   `json:"id"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	ProjectId string
}
