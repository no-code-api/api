package services

import (
	"github.com/no-code-api/api/internal/users/application/requests"
	"github.com/no-code-api/api/internal/users/application/responses"
)

type IService interface {
	Login(request *requests.LoginRequest) (*responses.LoginResponse, error)
	FindAll() ([]responses.UserResponse, error)
	FindById(id uint) (*responses.UserResponse, error)
	Create(request *requests.CreateUserRequest) error
	Update(request *requests.UpdateUserRequest) error
	DeleteById(id uint) error
}
