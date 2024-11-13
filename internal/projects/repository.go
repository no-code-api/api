package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

type IProjectRepository interface {
	Create(project *Project) bool
	FindByUser(userId uint) ([]*Project, bool)
	FindById(id string) (*Project, bool)
	Update(project *Project) bool
	DeleteById(id string) bool
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

func (r *projectRepository) Create(project *Project) bool {
	project.Id = generateUniqueId()
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

func (r *projectRepository) FindByUser(userId uint) ([]*Project, bool) {
	return r.findProjects(&findFilter{UserId: userId})
}

func (r *projectRepository) FindById(id string) (*Project, bool) {
	projects, ok := r.findProjects(&findFilter{Id: id})
	if !ok {
		return nil, false
	}
	var project *Project = nil
	if len(projects) > 0 {
		project = projects[0]
	}
	return project, true
}

func (r *projectRepository) Update(project *Project) bool {
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

func (r *projectRepository) findProjects(filter *findFilter) ([]*Project, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine(r.getQuery())
	query.AppendLine(r.getQueryFilter(filter))
	result, err := r.connection.ExecuteQuery(query.String())
	if err != nil {
		return nil, false
	}

	var projects []*Project
	for result.Next() {
		user := &Project{
			Id:          result.ReadString("id"),
			Name:        result.ReadString("name"),
			Description: result.ReadString("description"),
		}
		projects = append(projects, user)
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

func (r *projectRepository) getQueryFilter(filter *findFilter) string {
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
