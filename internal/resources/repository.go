package resources

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

type IRepository interface {
	CreateResource(resource *Resource) bool
	FindById(id string) (*Resource, bool)
	UpdateResource(resource *Resource) bool
	CheckResourcePathAvailableByProject(projectId string, path string) (bool, bool)
	CheckResourcePathAvailableByResourceId(resourceId string, path string) (bool, bool)
	FindAllResource(projectId string) ([]*Resource, bool)
	DeleteById(id string) bool
}

type repository struct {
	connection *postgre.Connection
	logger     *logger.Logger
}

func NewRepository(connection *postgre.Connection) IRepository {
	return &repository{
		connection: connection,
		logger:     logger.NewLogger("ResourcesRepository"),
	}
}

func (r *repository) CreateResource(resource *Resource) bool {
	resource.Id = core.GenerateUniqueId()
	if ok := r.addResource(resource); !ok {
		return false
	}
	if ok := r.upsertEndpoints(resource); !ok {
		return false
	}
	return true
}

func (r *repository) UpdateResource(resource *Resource) bool {
	if ok := r.updateResource(resource); !ok {
		return false
	}
	if ok := r.upsertEndpoints(resource); !ok {
		return false
	}
	return true
}

func (r *repository) FindAllResource(projectId string) ([]*Resource, bool) {
	return r.findResources(&findResourceFilter{ProjectId: projectId})
}

func (r *repository) FindById(id string) (*Resource, bool) {
	resources, ok := r.findResources(&findResourceFilter{Id: id})
	if !ok {
		return nil, false
	}
	var resource *Resource = nil
	if len(resources) > 0 {
		resource = resources[0]
	}
	return resource, true
}

func (r *repository) CheckResourcePathAvailableByProject(projectId string, path string) (bool, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine("SELECT COUNT(0)")
	query.AppendLine("FROM resources")
	query.AppendFormat("WHERE projectId=%s", utils.SqlString(projectId)).AppendNewLine()
	query.AppendFormat("AND path=%s", utils.SqlString(path)).AppendNewLine()
	count, err := r.connection.ExecuteSingleQuery(query.String())
	if err != nil {
		return false, false
	}
	return count.(int64) <= 0, true
}

func (r *repository) CheckResourcePathAvailableByResourceId(resourceId string, path string) (bool, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine("SELECT COUNT(0)")
	query.AppendLine("FROM resources")
	query.AppendFormat("WHERE id<>%s", utils.SqlString(resourceId)).AppendNewLine()
	query.AppendFormat("AND path=%s", utils.SqlString(path)).AppendNewLine()
	count, err := r.connection.ExecuteSingleQuery(query.String())
	if err != nil {
		return false, false
	}
	return count.(int64) <= 0, true
}

func (r *repository) DeleteById(id string) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("DELETE FROM resources")
	command.AppendFormat("WHERE id=%s", utils.SqlString(id))
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to delete resource: %s", err.Error())
		return false
	}
	return true
}

func (r *repository) addResource(resource *Resource) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("INSERT INTO resources")
	command.AppendLine("(id, projectId, path, createdAt, updatedAt)")
	command.AppendFormat("VALUES (%s", utils.SqlString(resource.Id)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(resource.ProjectId)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(resource.Path)).AppendNewLine()
	command.AppendLine(",NOW()")
	command.AppendLine(",NOW())")
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to insert resource: %s", err.Error())
		return false
	}
	return true
}

func (r *repository) upsertEndpoints(resource *Resource) bool {
	for _, endpoint := range resource.Endpoints {
		if endpoint.Id != 0 {
			if ok := r.updateEndpoint(endpoint); !ok {
				return false
			}
		} else {
			if ok := r.addEndpoint(resource.Id, endpoint); !ok {
				return false
			}
		}
	}
	return true
}

func (r *repository) addEndpoint(resourceId string, endpoint *Endpoint) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("INSERT INTO endpoints")
	command.AppendLine("(path, method, resourceId, createdAt, updatedAt)")
	command.AppendFormat("VALUES (%s", utils.SqlString(endpoint.Path)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(endpoint.Method)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(resourceId)).AppendNewLine()
	command.AppendLine(",NOW()")
	command.AppendLine(",NOW())")
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to insert endpoint: %s", err.Error())
		return false
	}
	return true
}

func (r *repository) updateResource(resource *Resource) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("UPDATE resources")
	command.AppendFormat("SET path=%s", utils.SqlString(resource.Path)).AppendNewLine()
	command.AppendLine(",updatedAt=NOW()")
	command.AppendFormat("WHERE id=%s", utils.SqlString(resource.Id)).AppendNewLine()
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to update resource: %s", err.Error())
		return false
	}
	return true
}

func (r *repository) updateEndpoint(endpoint *Endpoint) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("UPDATE endpoints")
	command.AppendFormat("SET path=%s", utils.SqlString(endpoint.Path)).AppendNewLine()
	command.AppendFormat(",method=%s", utils.SqlString(endpoint.Method)).AppendNewLine()
	command.AppendLine(",updatedAt=NOW()")
	command.AppendFormat("WHERE id=%d", endpoint.Id).AppendNewLine()
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logger.ErrorF("error to update endpoint: %s", err.Error())
		return false
	}
	return true
}

func (r *repository) findResources(filter *findResourceFilter) ([]*Resource, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine(r.getQuery())
	query.AppendLine(r.getQueryFilter(filter))
	result, err := r.connection.ExecuteQuery(query.String())
	if err != nil {
		return nil, false
	}

	var resources []*Resource
	var resourcesMap map[string]*Resource = make(map[string]*Resource)
	for result.Next() {
		resourceId := result.ReadString("resourceid")
		resource, exists := resourcesMap[resourceId]
		if !exists {
			resource = &Resource{
				Id:        resourceId,
				Path:      result.ReadString("resourcepath"),
				Endpoints: make([]*Endpoint, 0),
			}
			resourcesMap[resourceId] = resource
			resources = append(resources, resource)
		}
		endpoint := &Endpoint{
			Id:     uint(result.ReadInt("endpointid")),
			Path:   result.ReadString("endpointpath"),
			Method: result.ReadString("endpointmethod"),
		}
		resource.Endpoints = append(resource.Endpoints, endpoint)
	}
	return resources, true
}

func (r *repository) getQuery() string {
	return utils.NewStringBuilder().
		AppendLine("SELECT r.id resourceid").
		AppendLine(",r.path AS resourcepath").
		AppendLine(",e.id endpointid").
		AppendLine(",e.path AS endpointpath").
		AppendLine(",e.method AS endpointmethod").
		AppendLine("FROM resources AS r").
		AppendLine("LEFT JOIN endpoints AS e").
		AppendLine("ON e.resourceId=r.id").
		String()
}

func (r *repository) getQueryFilter(filter *findResourceFilter) string {
	query := utils.NewStringBuilder()
	query.AppendLine("WHERE 1=1")
	if filter.Id != "" {
		query.AppendFormat("AND r.id=%s", utils.SqlString(filter.Id))
	}
	if filter.ProjectId != "" {
		query.AppendFormat("AND r.projectId=%s", utils.SqlString(filter.ProjectId))
	}
	return query.String()
}
