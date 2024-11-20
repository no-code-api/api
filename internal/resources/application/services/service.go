package services

import (
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/internal/resources/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/resources/application/responses"
	dataRep "github.com/leandro-d-santos/no-code-api/internal/resources/data/repositories"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/models"
	domainRep "github.com/leandro-d-santos/no-code-api/internal/resources/domain/repositories"
	"github.com/leandro-d-santos/no-code-api/internal/resources/domain/validations"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type resourceService struct {
	resourceRepository domainRep.IRepository
	projectRepository  projects.IProjectRepository
}

func NewService(connection *postgre.Connection) IService {
	return resourceService{
		resourceRepository: dataRep.NewRepository(connection),
		projectRepository:  projects.NewRepository(connection),
	}
}

func (s resourceService) Create(createResource *requests.CreateResourceRequest) error {
	if _, err := s.findProject(createResource.ProjectId); err != nil {
		return err
	}
	if err := s.resourcePathAvailableByProject(createResource.ProjectId, createResource.Path); err != nil {
		return err
	}
	resource := createResource.ToModel()
	if err := validations.CreateResourceIsValid(resource); err != nil {
		return err
	}
	if ok := s.resourceRepository.CreateResource(resource); !ok {
		return errors.New("erro ao cadastrar recurso")
	}
	return nil
}

func (s resourceService) FindAll(projectId string) ([]responses.FindResourceResponse, error) {
	if _, err := s.findProject(projectId); err != nil {
		return nil, err
	}
	resources, ok := s.resourceRepository.FindAllResource(projectId)
	if !ok {
		return nil, errors.New("erro ao consultar recursos")
	}

	resourcesReponse := make([]responses.FindResourceResponse, len(resources))
	for index, resource := range resources {
		response := responses.FindResourceResponse{}
		response.FromModel(resource)
		resourcesReponse[index] = response
	}
	return resourcesReponse, nil
}

func (s resourceService) Update(updateResource *requests.UpdateResourceRequest) error {
	resource, err := s.findResourceById(updateResource.Id)
	if err != nil {
		return err
	}

	if err := s.resourcePathAvailableByResource(updateResource.Id, updateResource.Path); err != nil {
		return err
	}

	resource.Path = updateResource.Path
	resource.Endpoints = s.transformEndpointsRequestToModel(updateResource.Endpoints)
	if err := validations.UpdateResourceIsValid(resource); err != nil {
		return err
	}

	if ok := s.resourceRepository.UpdateResource(resource); !ok {
		return errors.New("erro ao atualizar recurso")
	}

	return nil
}

func (s resourceService) DeleteById(id string) error {
	if _, err := s.findResourceById(id); err != nil {
		return err
	}

	if ok := s.resourceRepository.DeleteById(id); !ok {
		return errors.New("erro ao remover recurso")
	}
	return nil
}

func (s resourceService) findProject(projectId string) (*projects.Project, error) {
	project, ok := s.projectRepository.FindById(projectId)
	if !ok {
		return nil, errors.New("erro ao consultar projeto")
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%v' não encontrado.", projectId)
		return nil, errors.New(message)
	}
	return project, nil
}

func (s resourceService) resourcePathAvailableByProject(projectId string, path string) error {
	available, ok := s.resourceRepository.CheckResourcePathAvailableByProject(projectId, path)
	if !ok {
		return errors.New("erro ao consultar disponibilidade de recurso")
	}
	if !available {
		message := fmt.Sprintf("Recurso '%s' não disponível", path)
		return errors.New(message)
	}
	return nil
}

func (s resourceService) resourcePathAvailableByResource(resourceId string, path string) error {
	available, ok := s.resourceRepository.CheckResourcePathAvailableByResourceId(resourceId, path)
	if !ok {
		return errors.New("erro ao consultar disponibilidade de recurso")
	}
	if !available {
		message := fmt.Sprintf("Recurso '%s' não disponível", path)
		return errors.New(message)
	}
	return nil
}

func (s resourceService) findResourceById(id string) (*models.Resource, error) {
	resource, ok := s.resourceRepository.FindById(id)
	if !ok {
		return nil, errors.New("erro ao consultar recurso")
	}
	if resource == nil {
		message := fmt.Sprintf("Recurso '%v' não encontrado.", id)
		return nil, errors.New(message)
	}
	return resource, nil
}

func (s resourceService) transformEndpointsRequestToModel(requestEndpoints []*requests.UpdateEndpointRequest) []*models.Endpoint {
	endpoints := make([]*models.Endpoint, len(requestEndpoints))
	for i, requestEndpoint := range requestEndpoints {
		endpoints[i] = requestEndpoint.ToModel()
	}
	return endpoints
}

func (s resourceService) findEndpointById(id string, endpoints []*models.Endpoint) *models.Endpoint {
	for _, endpoint := range endpoints {
		if endpoint.Id == id {
			return endpoint
		}
	}
	return nil
}
