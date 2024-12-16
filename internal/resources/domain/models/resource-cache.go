package models

type ResourceCache struct {
	Exists    bool
	Path      string           `json:"path"`
	Endpoints []*EndpointCache `json:"endpoints"`
}
