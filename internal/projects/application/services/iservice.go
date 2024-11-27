package services

import (
	"github.com/no-code-api/no-code-api/internal/projects/application/requests"
	"github.com/no-code-api/no-code-api/internal/projects/application/responses"
)

type IService interface {
	Create(request *requests.CreateProjectRequest) error
	FindByUser(userId uint) ([]responses.ProjectResponse, error)
	Update(request *requests.UpdateProjectRequest) error
	DeleteById(id string) error
}
