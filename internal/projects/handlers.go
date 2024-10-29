package projects

import (
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type ProjectHandler struct {
	DefaultPath    string
	projectService ProjectService
}

func NewHandler() ProjectHandler {
	return ProjectHandler{
		DefaultPath:    "/projects",
		projectService: NewService(),
	}
}

func (handler ProjectHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	userId, ok := baseHandler.GetUserId()
	if !ok {
		return
	}
	project := &CreateProjectViewModel{}
	if !baseHandler.BindJson(project) {
		return
	}
	project.UserId = userId
	if err := handler.projectService.Create(project); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData("Projeto criado com sucesso.")
}

func (handler ProjectHandler) HandleFindByUser(h *handler.BaseHandler) {
	userId, ok := h.GetUserId()
	if !ok {
		return
	}
	projects, err := handler.projectService.FindByUser(userId)
	if err != nil {
		h.BadRequest(err.Error())
		return
	}
	h.OkData(projects)
}

func (handler ProjectHandler) HandleUpdate(baseHandler *handler.BaseHandler) {
	id := baseHandler.Param("projectId")
	if id == "" {
		baseHandler.BadRequest("Código projeto não informado")
		return
	}
	project := &UpdateProjectViewModel{}
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
	if err := handler.projectService.Delete(id); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkMessage("Projeto deletado com sucesso.")
}
