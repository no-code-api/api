package endpoints

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/projects"
)

type findEndpointFilter struct {
	ProjectId string
	Id        uint
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
