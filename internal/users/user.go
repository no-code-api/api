package users

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
)

type filter struct {
	Id    uint
	Email string
}

type User struct {
	core.Entity
	Id       uint   `gorm:"unique;primaryKey;autoIncrement"`
	Name     string `gorm:"size:150;notnull"`
	Email    string `gorm:"size:100;unique;notnull"`
	Password string `gorm:"size:60;notnull"`
}
