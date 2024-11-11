package endpoints

import (
	"github.com/leandro-d-santos/no-code-api/internal/projects"
)

type findEndpointFilter struct {
	ProjectId string
	Id        uint
}

type Endpoint struct {
	Id        uint
	Path      string
	Method    string
	Schema    string
	ProjectId string
	Project   projects.Project
}
