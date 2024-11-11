package resources

import (
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

type ResourceService struct {
	resourceRepository IRepository
	projectRepository  projects.IProjectRepository
}

func NewService() ResourceService {
	connection := database.GetConnection()
	return ResourceService{
		resourceRepository: NewRepository(connection),
		projectRepository:  projects.NewRepository(connection),
	}
}

func (s ResourceService) Create(createResource *CreateResourceRequest) error {
	project, err := s.findProject(createResource.ProjectId)
	if err != nil {
		return err
	}

	resource := createResource.ToModel()
	resource.Project = *project

	if err := s.resourcePathAvailable(resource); err != nil {
		return err
	}

	if ok := s.resourceRepository.CreateResource(resource); !ok {
		return errors.New("erro ao cadastrar endpoint")
	}

	return nil
}

func (s ResourceService) FindAll(projectId string) ([]FindResourceResponse, error) {
	if _, err := s.findProject(projectId); err != nil {
		return nil, err
	}
	resources, ok := s.resourceRepository.FindAllResource(projectId)
	if !ok {
		return nil, errors.New("Erro ao consultar recursos")
	}

	resourcesReponse := make([]FindResourceResponse, len(resources))
	for index, resource := range resources {
		response := FindResourceResponse{}
		response.FromModel(resource)
		resourcesReponse[index] = response
	}
	return resourcesReponse, nil
}

func (s ResourceService) Update(ev *UpdateEndpointRequest) error {
	if _, err := s.findProject(ev.ProjectId); err != nil {
		return err
	}
	endpoint, err := s.findEndpoint(ev.ProjectId, ev.Id)
	if err != nil {
		return err
	}

	endpoint.Path = ev.Path
	endpoint.Method = ev.Method

	// if err := s.resourcePathAvailable(endpoint); err != nil {
	// 	return err
	// }

	// if ok := s.resourceRepository.UpdateResource(endpoint); !ok {
	// 	return errors.New("Erro ao atualizar endpoint.")
	// }

	return nil
}

func (s ResourceService) Delete(projectId string, endpointId uint) error {
	if _, err := s.findProject(projectId); err != nil {
		return err
	}

	if _, err := s.findEndpoint(projectId, endpointId); err != nil {
		return err
	}

	if ok := s.resourceRepository.DeleteEndpoint(projectId, endpointId); !ok {
		return errors.New("Erro ao cadastrar endpoint.")
	}
	return nil
}

func (s ResourceService) findProject(projectId string) (*projects.Project, error) {
	project, ok := s.projectRepository.FindById(projectId)
	if !ok {
		return nil, errors.New("Erro ao consultar projeto")
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%v' não encontrado.", projectId)
		return nil, errors.New(message)
	}
	return project, nil
}

func (s ResourceService) resourcePathAvailable(resource *Resource) error {
	available, ok := s.resourceRepository.ResourcePathAvailable(resource.ProjectId, resource.Path)
	if !ok {
		return errors.New("Erro ao consultar disponibilidade de recurso.")
	}
	if !available {
		message := fmt.Sprintf("Recurso '%s' não disponível", resource.Path)
		return errors.New(message)
	}
	return nil
}

func (s ResourceService) findEndpoint(projectId string, endpointId uint) (*Endpoint, error) {
	endpoint, ok := s.resourceRepository.FindEndpointById(projectId, endpointId)
	if !ok {
		return nil, errors.New("Erro ao consultar endpoint.")
	}
	if endpoint == nil {
		message := fmt.Sprintf("Endpoint '%v' não encontrado.", endpointId)
		return nil, errors.New(message)
	}
	return endpoint, nil
}
