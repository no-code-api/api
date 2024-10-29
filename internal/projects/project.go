package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

type findFilter struct {
	Id     string
	UserId uint
}

type Project struct {
	core.Entity
	Id     string     `gorm:"size:32;primaryKey"`
	UserId uint       `gorm:"notnull"`
	User   users.User `gorm:"foreignKey:UserId;references:Id"`
	Name   string     `gorm:"size:30;notnull"`
}
