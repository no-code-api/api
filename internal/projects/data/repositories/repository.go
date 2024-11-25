package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	dataModels "github.com/leandro-d-santos/no-code-api/internal/projects/data/models"
	domainModels "github.com/leandro-d-santos/no-code-api/internal/projects/domain/models"
	"github.com/leandro-d-santos/no-code-api/internal/projects/domain/repositories"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

type projectRepository struct {
	connection *postgre.Connection
	logger     *logger.Logger
}

func NewRepository(connection *postgre.Connection) repositories.IRepository {
	return &projectRepository{
		connection: connection,
		logger:     logger.NewLogger("ProjectRepository"),
	}
}

func (r *projectRepository) Create(project *domainModels.Project) bool {
	project.Id = core.GenerateUniqueId()
	command := utils.NewStringBuilder()
	command.AppendLine("INSERT INTO projects")
	command.AppendLine("(id, userId, name, description, createdAt, updatedAt)")
	command.AppendFormat("VALUES (%s", utils.SqlString(project.Id)).AppendNewLine()
	command.AppendFormat(",%d", project.UserId).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(project.Name)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(project.Description)).AppendNewLine()
	command.AppendLine(",NOW()")
	command.AppendLine(",NOW())")
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to insert project: %s", err.Error())
		return false
	}
	return true
}

func (r *projectRepository) FindByUser(userId uint) ([]*domainModels.Project, bool) {
	return r.findProjects(&dataModels.FindFilter{UserId: userId})
}

func (r *projectRepository) FindById(id string) (*domainModels.Project, bool) {
	projects, ok := r.findProjects(&dataModels.FindFilter{Id: id})
	if !ok {
		return nil, false
	}
	var project *domainModels.Project = nil
	if len(projects) > 0 {
		project = projects[0]
	}
	return project, true
}

func (r *projectRepository) Update(project *domainModels.Project) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("UPDATE projects")
	command.AppendFormat("SET name=%s", utils.SqlString(project.Name)).AppendNewLine()
	command.AppendFormat(",description=%s", utils.SqlString(project.Description)).AppendNewLine()
	command.AppendLine(",updatedAt=NOW()")
	command.AppendFormat("WHERE id=%s", utils.SqlString(project.Id)).AppendNewLine()
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to update project: %s", err.Error())
		return false
	}
	return true
}

func (r *projectRepository) DeleteById(id string) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("DELETE FROM projects")
	command.AppendFormat("WHERE id=%s", utils.SqlString(id))
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to delete project: %s", err.Error())
		return false
	}
	return true
}

func (r *projectRepository) findProjects(filter *dataModels.FindFilter) ([]*domainModels.Project, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine(r.getQuery())
	query.AppendLine(r.getQueryFilter(filter))
	result, err := r.connection.ExecuteQuery(query.String())
	if err != nil {
		return nil, false
	}

	var projects []*domainModels.Project
	for result.Next() {
		project := &domainModels.Project{
			Id:          result.ReadString("id"),
			Name:        result.ReadString("name"),
			Description: result.ReadString("description"),
		}
		projects = append(projects, project)
	}
	return projects, true
}

func (r *projectRepository) getQuery() string {
	query := utils.NewStringBuilder()
	query.AppendLine("SELECT id")
	query.AppendLine(",name")
	query.AppendLine(",description")
	query.AppendLine("FROM projects")
	return query.String()
}

func (r *projectRepository) getQueryFilter(filter *dataModels.FindFilter) string {
	query := utils.NewStringBuilder()
	query.AppendLine("WHERE 1=1")
	if filter.Id != "" {
		query.AppendFormat("AND id=%s", utils.SqlString(filter.Id)).AppendNewLine()
	}
	if filter.UserId > 0 {
		query.AppendFormat("AND userId=%d", filter.UserId)
	}
	return query.String()
}
