package resources

import (
	"github.com/leandro-d-santos/no-code-api/internal/core"
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

type IRepository interface {
	CreateResource(resource *Resource) (ok bool)
	FindEndpointById(projectId string, id uint) (endpoint *Endpoint, ok bool)
	UpdateResource(resource *Resource) (ok bool)
	ResourcePathAvailable(projectId string, path string) (available bool, ok bool)
	EndpointPathAvailable(endPoint *Endpoint) (available bool, ok bool)
	FindAllResource(projectId string) (endpoints []*Resource, ok bool)
	DeleteEndpoint(projectId string, id uint) (ok bool)
}

type repository struct {
	connection *database.Connection
	logger     *logger.Logger
}

func NewRepository(connection *database.Connection) IRepository {
	return &repository{
		connection: connection,
		logger:     logger.NewLogger("EndpointRepository"),
	}
}

func (r *repository) CreateResource(resource *Resource) (ok bool) {
	resource.Id = core.GenerateUniqueId()
	resource.SetCreatedAt()
	resource.SetUpdatedAt()
	return r.connection.Save(resource, false)
}

func (r *repository) UpdateResource(resource *Resource) (ok bool) {
	resource.SetUpdatedAt()
	return r.connection.Save(resource, true)
}

func (r *repository) FindEndpointById(projectId string, id uint) (endpoint *Endpoint, ok bool) {
	var result *Endpoint
	query := "project_id=? and id=?"
	if ok := r.connection.FindQuery(&result, query, projectId, id); !ok {
		return nil, false
	}
	if result.Id == 0 {
		result = nil
	}
	return result, true
}

func (r *repository) ResourcePathAvailable(projectId string, path string) (available bool, ok bool) {
	var result *Resource
	query := "project_id=? and path=?"
	if ok := r.connection.FindQuery(&result, query, projectId, path); !ok {
		return false, false
	}
	return result.Id == "", true
}

func (r *repository) EndpointPathAvailable(endpoint *Endpoint) (available bool, ok bool) {
	var result *Endpoint
	query := "resource_id=? and id <> ? and method=? and path=?"
	if ok := r.connection.FindQuery(&result, query, endpoint.ResourceId, endpoint.Id, endpoint.Method, endpoint.Path); !ok {
		return false, false
	}
	return result.Id == 0, true
}

func (r *repository) FindAllResource(projectId string) (resources []*Resource, ok bool) {

	query := core.NewStringBuilder()
	query.AppendLine("SELECT r.id resourceId")
	query.AppendLine(",r.path AS resourcePath")
	query.AppendLine(",e.id endpointId")
	query.AppendLine(",e.path AS endpointPath")
	query.AppendLine("FROM public.resources AS r")
	query.AppendLine("LEFT JOIN public.endpoints AS e")
	query.AppendLine("ON e.resource_id=r.id")
	query.AppendFormat("WHERE project_id=%s", projectId)
	// result, ok := r.connection.Execute(query.String())
	// if !ok {
	// 	return nil, false
	// }

	// endpoints, ok := r.findEndpointsByResourceId()
	// if ok := r.connection.FindQuery(&resources, query, projectId); !ok {
	// 	return nil, false
	// }

	return resources, true
}

func (r *repository) DeleteEndpoint(projectId string, id uint) (ok bool) {
	filter := struct {
		ProjectId string
		Id        uint
	}{ProjectId: projectId, Id: id}
	return r.connection.Delete(&Endpoint{}, filter)
}

func (r *repository) findEndpointsByResourceId(resourceId string) (endPoints []*Endpoint, ok bool) {
	query := "resource_id=?"
	if ok := r.connection.FindQuery(&endPoints, query, resourceId); !ok {
		return nil, false
	}
	return endPoints, true
}
