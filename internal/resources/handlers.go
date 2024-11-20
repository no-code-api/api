package resources

import (
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type ResourceHandler struct {
	DefaultPath     string
	resourceService ResourceService
}

func NewEndpointHandler() ResourceHandler {
	return ResourceHandler{
		DefaultPath:     "/projects/:projectId/resources",
		resourceService: NewService(),
	}
}

func (handler *ResourceHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	projectId := GetProjectId(baseHandler)
	if projectId == "" {
		return
	}
	resource := &CreateResourceRequest{}
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
	resource := &UpdateResourceRequest{}
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
	baseHandler.OkData("Recurso atualizado com sucesso.")
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
	baseHandler.OkData("Recurso deletado com sucesso.")
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
