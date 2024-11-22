package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/projects/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/projects/application/responses"
	dataRep "github.com/leandro-d-santos/no-code-api/internal/projects/data/repositories"
	"github.com/leandro-d-santos/no-code-api/internal/projects/domain/core"
	"github.com/leandro-d-santos/no-code-api/internal/projects/domain/models"
	domainRep "github.com/leandro-d-santos/no-code-api/internal/projects/domain/repositories"
	"github.com/leandro-d-santos/no-code-api/pkg/mongodb"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type projectService struct {
	projectRepository domainRep.IRepository
	mongoClient       *mongo.Database
}

func NewService(connection *postgre.Connection) IService {
	return projectService{
		projectRepository: dataRep.NewRepository(connection),
		mongoClient:       mongodb.GetConnection(),
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
	collectionName := core.GetCollectionName(project.Id)
	if err := s.mongoClient.CreateCollection(context.Background(), collectionName); err != nil {
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
	collectionName := core.GetCollectionName(id)
	if err := s.mongoClient.Collection(collectionName).Drop(context.Background()); err != nil {
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
