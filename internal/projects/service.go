package projects

import (
	"errors"
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/users"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type ProjectService struct {
	projectRepository IProjectRepository
}

func NewService() ProjectService {
	return ProjectService{
		projectRepository: NewRepository(postgre.GetConnection()),
	}
}

func (s ProjectService) Create(pv *CreateProjectViewModel) error {
	project := &Project{
		Name: pv.Name,
		User: users.User{Id: pv.UserId},
	}
	if ok := s.projectRepository.Create(project); !ok {
		return errors.New("erro ao cadastrar projeto")
	}
	return nil
}

func (s ProjectService) FindByUser(userId uint) ([]ProjectResponseViewModel, error) {
	projects, ok := s.projectRepository.FindByUser(userId)
	if !ok {
		return nil, errors.New("erro ao consultar projetos")
	}
	projectsResponse := make([]ProjectResponseViewModel, len(projects))
	for index, project := range projects {
		response := ProjectResponseViewModel{}
		response.FromModel(project)
		projectsResponse[index] = response
	}
	return projectsResponse, nil
}

func (s ProjectService) Update(pv *UpdateProjectViewModel) error {
	project, err := s.findById(pv.Id)
	if err != nil {
		return err
	}
	project.Name = pv.Name
	if ok := s.projectRepository.Update(project); !ok {
		return errors.New("erro ao atualizar projeto")
	}
	return nil
}

func (s ProjectService) Delete(id string) error {
	_, err := s.findById(id)
	if err != nil {
		return err
	}
	if ok := s.projectRepository.DeleteById(id); !ok {
		return errors.New("erro ao remover projeto")
	}
	return nil
}

func (s ProjectService) findById(id string) (*Project, error) {
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
