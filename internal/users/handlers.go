package users

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
)

func HandleFindAll(c *gin.Context) {
	repository := NewRepository()
	users, err := repository.FindAll()
	if err != nil {
		utils.ResBadRequest(c, "Erro ao consultar usuários.")
		return
	}
	utils.ResOkData(c, users)
}

func HandleFindById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	repository := NewRepository()
	user, err := repository.FindById(uint(id))
	if err != nil {
		utils.ResBadRequest(c, "Erro ao consultar usuário.")
		return
	}
	if user == nil {
		utils.ResNotFound(c, "Usuário não encontrado.")
		return
	}
	utils.ResOkData(c, user)
}

func HandleCreate(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResBadRequest(c, "Entrada de dados inválida.")
		return
	}
	repository := NewRepository()
	err := repository.Create(&user)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao criar usuário.")
		return
	}
	utils.ResNoContent(c)
}

func HandleUpdate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	var requestUser User
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		utils.ResBadRequest(c, "Entrada de dados inválida.")
		return
	}
	if id != int64(requestUser.Id) {
		utils.ResInvalidParam(c, "id")
		return
	}
	repository := NewRepository()
	user, err := repository.FindById(requestUser.Id)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao consultar usuário.")
		return
	}
	if user == nil {
		utils.ResNotFound(c, "Usuário não encontrado.")
		return
	}
	user.Name = requestUser.Name
	err = repository.Update(user)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao salvar usuário.")
		return
	}
	utils.ResNoContent(c)
}

func HandleDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	repository := NewRepository()
	err := repository.Delete(uint(id))
	if err != nil {
		utils.ResBadRequest(c, "Erro ao remover usuário.")
		return
	}
	utils.ResNoContent(c)
}
