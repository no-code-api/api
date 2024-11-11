package resources

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/projects"
)

type Resource struct {
	core.Entity
	Id        string           `gorm:"size:32;notnull;primaryKey"`
	ProjectId string           `gorm:"size:32;notnull"`
	Path      string           `gorm:"size:50;notnull"`
	Project   projects.Project `gorm:"foreignKey:ProjectId;references:id"`
	Endpoints []*Endpoint      `gorm:"foreignKey:ResourceId"`
}

type Endpoint struct {
	core.Entity
	Id         uint      `gorm:"primaryKey;autoIncrement"`
	Path       string    `gorm:"size:50;notnull"`
	Method     string    `gorm:"size:10;notnull"`
	ResourceId string    `gorm:"size:32;notnull"`
	Resource   *Resource `gorm:"foreignKey:ResourceId;references:id"`
}
