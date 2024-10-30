package endpoints

import (
	"strconv"

	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type EndpointHandler struct {
	DefaultPath     string
	endpointService EndpointService
}

func NewEndpointHandler() EndpointHandler {
	return EndpointHandler{
		DefaultPath:     "/projects/:projectId/endpoints",
		endpointService: NewService(),
	}
}

func (handler *EndpointHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	endpoint := &CreateEndpointRequest{}
	if !baseHandler.BindJson(endpoint) {
		return
	}
	endpoint.ProjectId = projectId
	if err := handler.endpointService.Create(endpoint); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData("Endpoint criado com sucesso.")
}

func (handler *EndpointHandler) HandleFindAll(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}

	endPoints, err := handler.endpointService.FindAll(projectId)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}

	baseHandler.OkData(endPoints)
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

	endpoint := &UpdateEndpointRequest{}
	if !baseHandler.BindJson(endpoint) {
		return
	}
	if endpointId != endpoint.Id {
		baseHandler.InvalidParam("Código endpoint")
		return
	}
	endpoint.ProjectId = projectId
	if err := handler.endpointService.Update(endpoint); err != nil {
		baseHandler.BadRequest(err.Error())
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
	if err := handler.endpointService.Delete(projectId, endpointId); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData("Endpoint deletado com sucesso.")
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
