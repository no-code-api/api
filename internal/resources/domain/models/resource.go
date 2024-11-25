package models

type Resource struct {
	Id        string
	ProjectId string
	Path      string
	Endpoints []*Endpoint
}
