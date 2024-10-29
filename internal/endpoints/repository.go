package endpoints

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

type IRepository interface {
	CreateEndpoint(endpoint *Endpoint) (ok bool)
	FindEndpointById(projectId string, id uint) (endpoint *Endpoint, ok bool)
	UpdateEndpoint(endpoint *Endpoint) (ok bool)
	PathAvailable(endPoint *Endpoint) (available bool, ok bool)
	FindAllEndpoints(projectId string) (endpoints []*Endpoint, ok bool)
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

func (r *repository) CreateEndpoint(endpoint *Endpoint) (ok bool) {
	endpoint.Id = 0
	endpoint.SetCreatedAt()
	endpoint.SetUpdatedAt()
	return r.connection.Save(endpoint, false)
}

func (r *repository) UpdateEndpoint(endpoint *Endpoint) (ok bool) {
	endpoint.SetUpdatedAt()
	return r.connection.Save(endpoint, true)
}

func (r *repository) FindEndpointById(projectId string, id uint) (endpoint *Endpoint, ok bool) {
	var result *Endpoint
	filter := &findEndpointFilter{ProjectId: projectId, Id: id}
	if ok := r.connection.Find(&result, filter); !ok {
		return nil, false
	}
	if result.Id == 0 {
		result = nil
	}
	return result, true
}

func (r *repository) PathAvailable(endpoint *Endpoint) (available bool, ok bool) {
	var result *Endpoint
	query := "project_id=? and id <> ? and method=? and path=?"
	if ok := r.connection.FindQuery(&result, query, endpoint.ProjectId, endpoint.Id, endpoint.Method, endpoint.Path); !ok {
		return false, false
	}
	return result.Id == 0, true
}

func (r *repository) FindAllEndpoints(projectId string) (endpoints []*Endpoint, ok bool) {
	var result []*Endpoint
	filter := &findEndpointFilter{ProjectId: projectId}
	if ok := r.connection.Find(&result, filter); !ok {
		return nil, false
	}
	return result, true
}

func (r *repository) DeleteEndpoint(projectId string, id uint) (ok bool) {
	filter := struct {
		ProjectId string
		Id        uint
	}{ProjectId: projectId, Id: id}
	return r.connection.Delete(&Endpoint{}, filter)
}
