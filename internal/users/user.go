package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty"`
}
