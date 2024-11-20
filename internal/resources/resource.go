package resources

type Resource struct {
	Id        string
	ProjectId string
	Path      string
	Endpoints []*Endpoint
}

type Endpoint struct {
	Id         uint
	Path       string
	Method     string
	ResourceId string
}
