package users

import (
	"gorm.io/gorm"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type updateUserRequest struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	gorm.Model
	Id       uint   `gorm:"unique;primaryKey;autoIncrement"`
	Name     string `gorm:"size:150;notnull"`
	Email    string `gorm:"size:100;unique;notnull"`
	Password string `gorm:"size:60;notnull"`
}

func (user *createUserRequest) ToModel() *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (userResponse *UserResponse) FromModel(user *User) {
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
}
