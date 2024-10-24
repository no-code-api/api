package projects

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/users"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
)

func resUserInvalidUser(c *gin.Context) {
	utils.ResBadRequest(c, "Usuário inválido.")
}

func handleCreate(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists || userId == "" {
		resUserInvalidUser(c)
		return
	}

	projectRequest := &createProjectRequest{}
	if !utils.BindJson(c, projectRequest) {
		return
	}

	repository := NewRepository()
	project := &Project{
		Name: projectRequest.Name,
		User: users.User{Id: userId.(uint)},
	}
	if ok := repository.Create(project); !ok {
		utils.ResBadRequest(c, "Erro ao cadastrar projeto.")
		return
	}
	utils.ResOkData(c, "Projeto criado com sucesso.")
}

func handleFindByUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists || userId == "" {
		resUserInvalidUser(c)
		return
	}

	repository := NewRepository()
	projects, ok := repository.FindByUser(userId.(uint))
	if !ok {
		utils.ResBadRequest(c, "Erro ao consultar projetos.")
		return
	}
	projectsResponse := make([]projectResponse, len(projects))
	for index, project := range projects {
		response := projectResponse{}
		response.FromModel(project)
		projectsResponse[index] = response
	}
	utils.ResOkData(c, projectsResponse)
}

func handleUpdate(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists || userId == "" {
		resUserInvalidUser(c)
		return
	}
	projectRequest := &updateProjectRequest{}
	if !utils.BindJson(c, projectRequest) {
		return
	}

	repository := NewRepository()
	project, ok := repository.FindById(projectRequest.Id)
	if !ok {
		utils.ResBadRequest(c, "Projeto não encontrado.")
		return
	}

	project.Name = projectRequest.Name
	if ok := repository.Update(project); !ok {
		utils.ResBadRequest(c, "Erro ao cadastrar projeto.")
		return
	}
	utils.ResOkData(c, "Projeto criado com sucesso.")
}

func handleDeleteByUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists || userId == "" {
		resUserInvalidUser(c)
		return
	}

	id := c.Param("id")
	if id == "" {
		utils.ResInvalidParam(c, "id")
		return
	}

	repository := NewRepository()
	project, ok := repository.FindById(id)
	if !ok {
		utils.ResBadRequest(c, "Erro ao consultar projetos.")
		return
	}
	if project == nil {
		message := fmt.Sprintf("Projeto '%v' não existe.", id)
		utils.ResBadRequest(c, message)
		return
	}
	if ok := repository.DeleteById(id); !ok {
		utils.ResBadRequest(c, "Erro ao remover projeto")
		return
	}
	utils.ResOkMessage(c, "Projeto deletado com sucesso.")
}
