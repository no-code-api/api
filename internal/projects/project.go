package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

type findFilter struct {
	Id     string
	UserId uint
}

type Project struct {
	Id     string
	UserId uint
	User   users.User
	Name   string
}
