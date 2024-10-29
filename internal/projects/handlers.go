package projects

import (
	"fmt"

	"github.com/leandro-d-santos/no-code-api/internal/handler"
	"github.com/leandro-d-santos/no-code-api/internal/users"
)

func handleCreate(h *handler.BaseHandler) {
	userId, ok := h.GetUserId()
	if !ok {
		return
	}

	projectRequest := &createProjectRequest{}
	if !h.BindJson(projectRequest) {
		return
	}

	repository := NewRepository()
	project := &Project{
		Name: projectRequest.Name,
		User: users.User{Id: userId},
	}
	if ok := repository.Create(project); !ok {
		h.BadRequest("Erro ao cadastrar projeto.")
		return
	}
	h.OkData("Projeto criado com sucesso.")
}

func handleFindByUser(h *handler.BaseHandler) {
	userId, ok := h.GetUserId()
	if !ok {
		return
	}

	repository := NewRepository()
	projects, ok := repository.FindByUser(userId)
	if !ok {
		h.BadRequest("Erro ao consultar projetos.")
		return
	}
	projectsResponse := make([]projectResponse, len(projects))
	for index, project := range projects {
		response := projectResponse{}
		response.FromModel(project)
		projectsResponse[index] = response
	}
	h.OkData(projectsResponse)
}

func handleUpdate(h *handler.BaseHandler) {
	projectRequest := &updateProjectRequest{}
	if !h.BindJson(projectRequest) {
		return
	}

	repository := NewRepository()
	project, ok := repository.FindById(projectRequest.Id)
	if !ok {
		h.BadRequest("Projeto não encontrado.")
		return
	}

	project.Name = projectRequest.Name
	if ok := repository.Update(project); !ok {
		h.BadRequest("Erro ao cadastrar projeto.")
		return
	}
	h.OkData("Projeto criado com sucesso.")
}

func handleDeleteByUser(h *handler.BaseHandler) {
	id := h.Param("projectId")
	if id == "" {
		h.InvalidParam("Código projeto não informado")
		return
	}

	repository := NewRepository()
	project, ok := repository.FindById(id)
	if !ok {
		h.BadRequest("Erro ao consultar projetos.")
		return
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%v' não existe.", id)
		h.BadRequest(message)
		return
	}
	if ok := repository.DeleteById(id); !ok {
		h.BadRequest("Erro ao remover projeto")
		return
	}
	h.OkMessage("Projeto deletado com sucesso.")
}
