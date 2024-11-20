package resources

import (
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type ResourceService struct {
	resourceRepository IRepository
	projectRepository  projects.IProjectRepository
}

func NewService() ResourceService {
	connection := postgre.GetConnection()
	return ResourceService{
		resourceRepository: NewRepository(connection),
		projectRepository:  projects.NewRepository(connection),
	}
}

func (s ResourceService) Create(createResource *CreateResourceRequest) error {
	if _, err := s.findProject(createResource.ProjectId); err != nil {
		return err
	}
	if err := s.resourcePathAvailableByProject(createResource.ProjectId, createResource.Path); err != nil {
		return err
	}
	resource := createResource.ToModel()
	if err := CreateResourceIsValid(resource); err != nil {
		return err
	}
	if ok := s.resourceRepository.CreateResource(resource); !ok {
		return errors.New("erro ao cadastrar recurso")
	}
	return nil
}

func (s ResourceService) FindAll(projectId string) ([]FindResourceResponse, error) {
	if _, err := s.findProject(projectId); err != nil {
		return nil, err
	}
	resources, ok := s.resourceRepository.FindAllResource(projectId)
	if !ok {
		return nil, errors.New("erro ao consultar recursos")
	}

	resourcesReponse := make([]FindResourceResponse, len(resources))
	for index, resource := range resources {
		response := FindResourceResponse{}
		response.FromModel(resource)
		resourcesReponse[index] = response
	}
	return resourcesReponse, nil
}

func (s ResourceService) Update(updateResource *UpdateResourceRequest) error {
	resource, err := s.findResourceById(updateResource.Id)
	if err != nil {
		return err
	}

	if err := s.resourcePathAvailableByResource(updateResource.Id, updateResource.Path); err != nil {
		return err
	}

	resource.Path = updateResource.Path
	endpoints := make([]*Endpoint, 0)
	endpoints = append(endpoints, resource.Endpoints...)
	s.populateResourceEndpoints(endpoints, updateResource.Endpoints)
	if err := UpdateResourceIsValid(resource); err != nil {
		return err
	}

	if ok := s.resourceRepository.UpdateResource(resource); !ok {
		return errors.New("erro ao atualizar recurso")
	}

	return nil
}

func (s ResourceService) DeleteById(id string) error {
	if _, err := s.findResourceById(id); err != nil {
		return err
	}

	if ok := s.resourceRepository.DeleteById(id); !ok {
		return errors.New("erro ao remover recurso")
	}
	return nil
}

func (s ResourceService) findProject(projectId string) (*projects.Project, error) {
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

func (s ResourceService) resourcePathAvailableByProject(projectId string, path string) error {
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

func (s ResourceService) resourcePathAvailableByResource(resourceId string, path string) error {
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

func (s ResourceService) findResourceById(id string) (*Resource, error) {
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

func (s ResourceService) populateResourceEndpoints(endpoints []*Endpoint, requestEndpoints []*UpdateEndpointRequest) {
	for _, requestEndpoint := range requestEndpoints {
		if requestEndpoint.Id == 0 {
			endpoints = append(endpoints, requestEndpoint.ToModel())
			continue
		}
		endpoint := s.findEndpointById(requestEndpoint.Id, endpoints)
		if endpoint == nil {
			continue
		}
		endpoint.Method = requestEndpoint.Method
		endpoint.Path = requestEndpoint.Path
	}
}

func (s ResourceService) findEndpointById(id uint, endpoints []*Endpoint) *Endpoint {
	for _, endpoint := range endpoints {
		if endpoint.Id == id {
			return endpoint
		}
	}
	return nil
}
