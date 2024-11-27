package services

import (
	"github.com/no-code-api/api/internal/resources/application/requests"
	"github.com/no-code-api/api/internal/resources/application/responses"
)

type IService interface {
	Create(createResource *requests.CreateResourceRequest) error
	FindAll(projectId string) ([]responses.FindResourceResponse, error)
	Update(updateResource *requests.UpdateResourceRequest) error
	DeleteById(id string) error
}
