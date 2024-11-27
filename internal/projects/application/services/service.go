package services

import (
	"errors"
	"fmt"

	"github.com/no-code-api/no-code-api/internal/projects/application/requests"
	"github.com/no-code-api/no-code-api/internal/projects/application/responses"
	dataRep "github.com/no-code-api/no-code-api/internal/projects/data/repositories"
	"github.com/no-code-api/no-code-api/internal/projects/domain/models"
	domainRep "github.com/no-code-api/no-code-api/internal/projects/domain/repositories"
	"github.com/no-code-api/no-code-api/internal/resources/domain/services"
	"github.com/no-code-api/no-code-api/pkg/postgre"
)

type projectService struct {
	projectRepository          domainRep.IRepository
	resourceDynamicDataService services.IResourceDynamicDataService
}

func NewService(connection *postgre.Connection) IService {
	return projectService{
		projectRepository:          dataRep.NewRepository(connection),
		resourceDynamicDataService: services.NewResourceDynamicDataService(),
	}
}

func (s projectService) Create(request *requests.CreateProjectRequest) error {
	project := &models.Project{
		Name:        request.Name,
		Description: request.Description,
		UserId:      request.UserId,
	}
	if ok := s.projectRepository.Create(project); !ok {
		return errors.New("erro ao cadastrar projeto")
	}
	if err := s.resourceDynamicDataService.CreateCollection(project.Id); err != nil {
		return fmt.Errorf("erro ao criar coleção do projeto")
	}
	return nil
}

func (s projectService) FindByUser(userId uint) ([]responses.ProjectResponse, error) {
	projects, ok := s.projectRepository.FindByUser(userId)
	if !ok {
		return nil, errors.New("erro ao consultar projetos")
	}
	projectsResponse := make([]responses.ProjectResponse, len(projects))
	for index, project := range projects {
		response := responses.ProjectResponse{}
		response.FromModel(project)
		projectsResponse[index] = response
	}
	return projectsResponse, nil
}

func (s projectService) Update(request *requests.UpdateProjectRequest) error {
	project, err := s.findById(request.Id)
	if err != nil {
		return err
	}
	project.Name = request.Name
	project.Description = request.Description
	if ok := s.projectRepository.Update(project); !ok {
		return errors.New("erro ao atualizar projeto")
	}
	return nil
}

func (s projectService) DeleteById(id string) error {
	_, err := s.findById(id)
	if err != nil {
		return err
	}
	if ok := s.projectRepository.DeleteById(id); !ok {
		return errors.New("erro ao remover projeto")
	}

	if err := s.resourceDynamicDataService.DropCollection(id); err != nil {
		return fmt.Errorf("erro ao deletar coleção do projeto")
	}
	return nil
}

func (s projectService) findById(id string) (*models.Project, error) {
	project, ok := s.projectRepository.FindById(id)
	if !ok {
		return nil, errors.New("erro ao consultar projetos")
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%s' não existe", id)
		return nil, errors.New(message)
	}
	return project, nil
}
