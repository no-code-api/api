package handlers

import (
	"github.com/no-code-api/api/internal/handler"
	"github.com/no-code-api/api/internal/resources/application/requests"
	"github.com/no-code-api/api/internal/resources/application/services"
	"github.com/no-code-api/api/pkg/postgre"
)

type ResourceHandler struct {
	DefaultPath     string
	resourceService services.IService
}

func NewEndpointHandler() ResourceHandler {
	connection := postgre.GetConnection()
	return ResourceHandler{
		DefaultPath:     "/projects/:projectId/resources",
		resourceService: services.NewService(connection),
	}
}

func (handler *ResourceHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	resource := &requests.CreateResourceRequest{}
	if !baseHandler.BindJson(resource) {
		return
	}
	resource.ProjectId = projectId
	if err := handler.resourceService.Create(resource); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Created()
}

func (handler *ResourceHandler) HandleFindAll(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}

	endPoints, err := handler.resourceService.FindAll(projectId)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}

	baseHandler.OkData(endPoints)
}

func (handler *ResourceHandler) HandleUpdate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	resourceId := GetResourceId(baseHandler)
	if resourceId == "" {
		return
	}
	resource := &requests.UpdateResourceRequest{}
	if !baseHandler.BindJson(resource) {
		return
	}
	if resourceId != resource.Id {
		baseHandler.InvalidParam("Código recurso")
		return
	}
	resource.ProjectId = projectId
	if err := handler.resourceService.Update(resource); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkMessage("Recurso atualizado com sucesso.")
}

func (handler *ResourceHandler) HandleDelete(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	resourceId := GetResourceId(baseHandler)
	if resourceId == "" {
		return
	}
	if err := handler.resourceService.DeleteById(resourceId); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkMessage("Recurso deletado com sucesso.")
}

func GetProjectId(baseHandler *handler.BaseHandler) string {
	projectId := baseHandler.Param("projectId")
	if projectId == "" {
		baseHandler.BadRequest("Código projeto não informado.")
		return ""
	}
	return projectId
}

func GetResourceId(baseHandler *handler.BaseHandler) string {
	resourceId := baseHandler.Param("resourceId")
	if resourceId == "" {
		baseHandler.BadRequest("Código recurso não informado.")
		return ""
	}
	return resourceId
}
