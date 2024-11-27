package requests

import "github.com/no-code-api/no-code-api/internal/users/domain/models"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *CreateUserRequest) ToModel() *models.User {
	return &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
