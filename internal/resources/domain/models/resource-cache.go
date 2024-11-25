package models

type ResourceCache struct {
	Path      string           `json:"path"`
	Endpoints []*EndpointCache `json:"endpoints"`
}
