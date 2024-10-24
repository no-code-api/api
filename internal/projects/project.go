package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

type createProjectRequest struct {
	Name string `json:"name"`
}

type updateProjectRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type findFilter struct {
	Id     string
	UserId uint
}

type Project struct {
	core.Entity
	Id     string     `gorm:"size:32;unique;primaryKet;autoIncrement"`
	UserId uint       `gorm:"notnull"`
	User   users.User `gorm:"foreignKey:UserId;references:Id"`
	Name   string     `gorm:"size:30;notnull"`
}

type projectResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (projectResponse *projectResponse) FromModel(project *Project) {
	projectResponse.Id = project.Id
	projectResponse.Name = project.Name
}
