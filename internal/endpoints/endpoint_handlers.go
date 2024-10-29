package endpoints

import (
	"fmt"
	"strconv"

	"github.com/leandro-d-santos/no-code-api/internal/handler"
	"github.com/leandro-d-santos/no-code-api/internal/projects"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
)

type EndpointHandler struct {
	DefaultPath        string
	EndpointRepository IRepository
	ProjectRepository  projects.IProjectRepository
}

func NewEndpointHandler() EndpointHandler {
	connection := database.GetConnection()
	return EndpointHandler{
		DefaultPath:        "/projects/:projectId/endpoints",
		EndpointRepository: NewRepository(connection),
		ProjectRepository:  projects.NewRepository(),
	}
}

func (handler *EndpointHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	project := FindProject(baseHandler, handler.ProjectRepository, projectId)
	if project == nil {
		return
	}

	endpointRequest := &createEndpointRequest{}
	if !baseHandler.BindJson(endpointRequest) {
		return
	}

	endpoint := endpointRequest.ToModel()
	endpoint.ProjectId = projectId
	endpoint.Project = *project

	if !PathAvailable(baseHandler, handler.EndpointRepository, endpoint) {
		return
	}

	if ok := handler.EndpointRepository.CreateEndpoint(endpoint); !ok {
		baseHandler.BadRequest("Erro ao cadastrar endpoint.")
		return
	}
	baseHandler.OkData("Endpoint criado com sucesso.")
}

func (handler *EndpointHandler) HandleFindAll(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}

	if project := FindProject(baseHandler, handler.ProjectRepository, projectId); project == nil {
		return
	}
	endpoints, ok := handler.EndpointRepository.FindAllEndpoints(projectId)
	if !ok {
		baseHandler.BadRequest("Erro ao consultar endpoints")
		return
	}

	endPointsReponse := make([]endpointResponse, len(endpoints))
	for index, endpoint := range endpoints {
		response := endpointResponse{}
		response.FromModel(endpoint)
		endPointsReponse[index] = response
	}
	baseHandler.OkData(endPointsReponse)
}

func (handler *EndpointHandler) HandleUpdate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	endpointId := GetEndpointId(baseHandler)
	if endpointId == 0 {
		return
	}

	endpointRequest := &updateEndpointRequest{}
	if !baseHandler.BindJson(endpointRequest) {
		return
	}

	if endpointId != endpointRequest.Id {
		baseHandler.InvalidParam("Código endpoint")
		return
	}

	if project := FindProject(baseHandler, handler.ProjectRepository, projectId); project == nil {
		return
	}
	endpoint := FindEndpoint(baseHandler, handler.EndpointRepository, projectId, endpointId)
	if endpoint == nil {
		return
	}

	endpoint.Path = endpointRequest.Path
	endpoint.Method = endpointRequest.Method

	if !PathAvailable(baseHandler, handler.EndpointRepository, endpoint) {
		return
	}

	if ok := handler.EndpointRepository.UpdateEndpoint(endpoint); !ok {
		baseHandler.BadRequest("Erro ao atualizar endpoint.")
		return
	}
	baseHandler.OkData("Endpoint atualizado com sucesso.")
}

func (handler *EndpointHandler) HandleDelete(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	endpointId := GetEndpointId(baseHandler)
	if endpointId == 0 {
		return
	}

	if project := FindProject(baseHandler, handler.ProjectRepository, projectId); project == nil {
		return
	}

	if endpoint := FindEndpoint(baseHandler, handler.EndpointRepository, projectId, endpointId); endpoint == nil {
		return
	}

	if ok := handler.EndpointRepository.DeleteEndpoint(projectId, endpointId); !ok {
		baseHandler.BadRequest("Erro ao cadastrar endpoint.")
		return
	}
	baseHandler.OkData("Endpoint deletado com sucesso.")
}

func PathAvailable(baseHandler *handler.BaseHandler, repository IRepository, endpoint *Endpoint) bool {
	available, ok := repository.PathAvailable(endpoint)
	if !ok {
		baseHandler.BadRequest("Erro ao consultar disponibilidade de endpoint.")
		return false
	}
	if !available {
		message := fmt.Sprintf("Endpoint '%s' para o método '%s' não disponível", endpoint.Path, endpoint.Method)
		baseHandler.BadRequest(message)
		return false
	}
	return true
}

func FindProject(baseHandler *handler.BaseHandler, projectRepository projects.IProjectRepository, projectId string) *projects.Project {
	project, ok := projectRepository.FindProjectById(projectId)
	if !ok {
		baseHandler.BadRequest("Erro ao consultar projeto")
		return nil
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%v' não encontrado.", projectId)
		baseHandler.BadRequest(message)
		return nil
	}
	return project
}

func FindEndpoint(baseHandler *handler.BaseHandler, repository IRepository, projectId string, endpointId uint) *Endpoint {
	endpoint, ok := repository.FindEndpointById(projectId, endpointId)
	if !ok {
		baseHandler.BadRequest("Erro ao consultar endpoint.")
		return nil
	}
	if endpoint == nil {
		message := fmt.Sprintf("Endpoint '%v' não encontrado.", endpointId)
		baseHandler.BadRequest(message)
		return nil
	}
	return endpoint
}

func GetProjectId(baseHandler *handler.BaseHandler) string {
	projectId := baseHandler.Param("projectId")
	if projectId == "" {
		baseHandler.BadRequest("Código projeto não informado.")
		return ""
	}
	return projectId
}

func GetEndpointId(baseHandler *handler.BaseHandler) uint {
	endpointId, _ := strconv.ParseInt(baseHandler.Param("endpointId"), 10, 32)
	if endpointId == 0 {
		baseHandler.BadRequest("Código endpoint não informado.")
		return 0
	}
	return uint(endpointId)
}
