package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type IProjectRepository interface {
	Create(project *Project) (ok bool)
	FindByUser(userId uint) (projects []*Project, ok bool)
	FindById(id string) (project *Project, ok bool)
	Update(project *Project) (ok bool)
	DeleteById(id string) (ok bool)
}

type projectRepository struct {
	connection *postgre.Connection
	logger     *logger.Logger
}

func NewRepository(connection *postgre.Connection) IProjectRepository {
	return &projectRepository{
		connection: connection,
		logger:     logger.NewLogger("ProjectRepository"),
	}
}

func (r *projectRepository) Create(project *Project) (ok bool) {
	// project.Id = generateUniqueId()
	// project.SetCreatedAt()
	// project.SetUpdatedAt()
	// return r.connection.Save(project, false)
	return true
}

func (r *projectRepository) FindByUser(userId uint) (projects []*Project, ok bool) {
	// var result []*Project
	// filter := &findFilter{UserId: userId}
	// if ok := r.connection.Find(&result, filter); !ok {
	// 	return nil, false
	// }
	// return result, true
	return nil, true
}

func (r *projectRepository) FindById(id string) (project *Project, ok bool) {
	// var result *Project
	// filter := &findFilter{Id: id}
	// if ok := r.connection.Find(&result, filter); !ok {
	// 	return nil, false
	// }
	// if result.Id == "" {
	// 	result = nil
	// }
	// return result, true
	return nil, true
}

func (r *projectRepository) Update(project *Project) (ok bool) {
	// project.SetUpdatedAt()
	// return r.connection.Save(project, true)
	return true
}

func (r *projectRepository) DeleteById(id string) (ok bool) {
	// filter := struct{ Id string }{Id: id}
	// return r.connection.Delete(&Project{}, filter)
	return true
}
