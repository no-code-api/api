package handlers

import (
	"github.com/leandro-d-santos/no-code-api/internal/handler"
	"github.com/leandro-d-santos/no-code-api/internal/projects/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/projects/application/services"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type ProjectHandler struct {
	DefaultPath    string
	projectService services.IService
}

func NewHandler() ProjectHandler {
	connection := postgre.GetConnection()
	return ProjectHandler{
		DefaultPath:    "/projects",
		projectService: services.NewService(connection),
	}
}

func (handler ProjectHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	userId, ok := baseHandler.GetUserId()
	if !ok {
		return
	}
	project := &requests.CreateProjectRequest{}
	if !baseHandler.BindJson(project) {
		return
	}
	project.UserId = userId
	if err := handler.projectService.Create(project); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Created()
}

func (handler ProjectHandler) HandleFindByUser(baseHandler *handler.BaseHandler) {
	userId, ok := baseHandler.GetUserId()
	if !ok {
		return
	}
	projects, err := handler.projectService.FindByUser(userId)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData(projects)
}

func (handler ProjectHandler) HandleUpdate(baseHandler *handler.BaseHandler) {
	id := baseHandler.Param("projectId")
	if id == "" {
		baseHandler.BadRequest("Código projeto não informado")
		return
	}
	project := &requests.UpdateProjectRequest{}
	if !baseHandler.BindJson(project) {
		return
	}
	if project.Id != id {
		baseHandler.InvalidParam("Código projeto")
		return
	}
	if err := handler.projectService.Update(project); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData("Projeto atualizado com sucesso.")
}

func (handler ProjectHandler) HandleDeleteByUser(baseHandler *handler.BaseHandler) {
	id := baseHandler.Param("projectId")
	if id == "" {
		baseHandler.BadRequest("Código projeto não informado")
		return
	}
	if err := handler.projectService.DeleteById(id); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkMessage("Projeto deletado com sucesso.")
}
