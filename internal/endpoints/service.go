package endpoints

import (
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

type EndpointService struct {
	endpointRepository IRepository
	projectRepository  projects.IProjectRepository
}

func NewService() EndpointService {
	connection := database.GetConnection()
	return EndpointService{
		endpointRepository: NewRepository(connection),
		projectRepository:  projects.NewRepository(connection),
	}
}

func (s EndpointService) Create(ev *CreateEndpointRequest) error {
	project, err := s.findProject(ev.ProjectId)
	if err != nil {
		return err
	}

	endpoint := ev.ToModel()
	endpoint.Project = *project

	if err := s.pathAvailable(endpoint); err != nil {
		return err
	}

	if ok := s.endpointRepository.CreateEndpoint(endpoint); !ok {
		return errors.New("Erro ao cadastrar endpoint.")
	}

	return nil
}

func (s EndpointService) FindAll(projectId string) ([]FindEndpointResponse, error) {
	if _, err := s.findProject(projectId); err != nil {
		return nil, err
	}
	endpoints, ok := s.endpointRepository.FindAllEndpoints(projectId)
	if !ok {
		return nil, errors.New("Erro ao consultar endpoints")
	}

	endpointsReponse := make([]FindEndpointResponse, len(endpoints))
	for index, endpoint := range endpoints {
		response := FindEndpointResponse{}
		response.FromModel(endpoint)
		endpointsReponse[index] = response
	}
	return endpointsReponse, nil
}

func (s EndpointService) Update(ev *UpdateEndpointRequest) error {
	if _, err := s.findProject(ev.ProjectId); err != nil {
		return err
	}
	endpoint, err := s.findEndpoint(ev.ProjectId, ev.Id)
	if err != nil {
		return err
	}

	endpoint.Path = ev.Path
	endpoint.Method = ev.Method

	if err := s.pathAvailable(endpoint); err != nil {
		return err
	}

	if ok := s.endpointRepository.UpdateEndpoint(endpoint); !ok {
		return errors.New("Erro ao atualizar endpoint.")
	}

	return nil
}

func (s EndpointService) Delete(projectId string, endpointId uint) error {
	if _, err := s.findProject(projectId); err != nil {
		return err
	}

	if _, err := s.findEndpoint(projectId, endpointId); err != nil {
		return err
	}

	if ok := s.endpointRepository.DeleteEndpoint(projectId, endpointId); !ok {
		return errors.New("Erro ao cadastrar endpoint.")
	}
	return nil
}

func (s EndpointService) findProject(projectId string) (*projects.Project, error) {
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

func (s EndpointService) pathAvailable(endpoint *Endpoint) error {
	available, ok := s.endpointRepository.PathAvailable(endpoint)
	if !ok {
		return errors.New("Erro ao consultar disponibilidade de endpoint.")
	}
	if !available {
		message := fmt.Sprintf("Endpoint '%s' para o método '%s' não disponível", endpoint.Path, endpoint.Method)
		return errors.New(message)
	}
	return nil
}

func (s EndpointService) findEndpoint(projectId string, endpointId uint) (*Endpoint, error) {
	endpoint, ok := s.endpointRepository.FindEndpointById(projectId, endpointId)
	if !ok {
		return nil, errors.New("Erro ao consultar endpoint.")
	}
	if endpoint == nil {
		message := fmt.Sprintf("Endpoint '%v' não encontrado.", endpointId)
		return nil, errors.New(message)
	}
	return endpoint, nil
}
