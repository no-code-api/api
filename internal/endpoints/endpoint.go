package endpoints

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/projects"
)

type createEndpointRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type updateEndpointRequest struct {
	Id     uint   `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type findEndpointFilter struct {
	ProjectId string
	Id        uint
}

type endpointResponse struct {
	Id     uint   `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type Endpoint struct {
	core.Entity
	Id        uint             `gorm:"primaryKey;autoIncrement"`
	Path      string           `gorm:"size:50;notnull"`
	Method    string           `gorm:"size:10;notnull"`
	Schema    string           `gorm:"size:300;notnull"`
	ProjectId string           `gorm:"size:32;notnull"`
	Project   projects.Project `gorm:"foreignKey:ProjectId;references:Id"`
}

func (ep *createEndpointRequest) ToModel() *Endpoint {
	return &Endpoint{
		Path:   ep.Path,
		Method: ep.Method,
	}
}

func (ep *endpointResponse) FromModel(endpoint *Endpoint) {
	ep.Id = endpoint.Id
	ep.Path = endpoint.Path
	ep.Method = endpoint.Method
}
