package users

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
	"github.com/leandro-d-santos/no-code-api/pkg/jwt"
)

func errorToSearchUser(c *gin.Context) {
	utils.ResBadRequest(c, "Erro ao consultar usuário.")
}

func HandleLogin(c *gin.Context) {
	userRequest := &loginRequest{}
	if err := utils.BindJson(c, userRequest); err != nil {
		return
	}

	repository := NewRepository()
	user, err := repository.FindByEmail(userRequest.Email)
	if err != nil {
		errorToSearchUser(c)
		return
	}

	if user == nil {
		utils.ResBadRequest(c, "Email inválido.")
		return
	}

	if !VerifyPassword(userRequest.Password, user.Password) {
		utils.ResBadRequest(c, "Senha inválida.")
		return
	}

	token, err := jwt.GenerateJWT(user.Id)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao gerar token.")
		return
	}

	response := &loginResponse{Token: token}
	utils.ResOkData(c, response)
}

func HandleFindAll(c *gin.Context) {
	repository := NewRepository()
	users, err := repository.FindAll()
	if err != nil {
		errorToSearchUser(c)
		return
	}
	usersReponse := make([]*UserResponse, len(users))
	for index, user := range users {
		userResponse := &UserResponse{}
		userResponse.FromModel(user)
		usersReponse[index] = userResponse
	}
	utils.ResOkData(c, usersReponse)
}

func HandleFindById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	repository := NewRepository()
	user, err := repository.FindById(uint(id))
	if err != nil {
		errorToSearchUser(c)
		return
	}
	if user == nil {
		utils.ResNotFound(c, "Usuário não encontrado.")
		return
	}
	userResponse := &UserResponse{}
	userResponse.FromModel(user)
	utils.ResOkData(c, userResponse)
}

func HandleCreate(c *gin.Context) {
	var user createUserRequest
	if err := utils.BindJson(c, &user); err != nil {
		return
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao gerar senha.")
	}

	user.Password = hash
	repository := NewRepository()
	if err := repository.Create(user.ToModel()); err != nil {
		utils.ResBadRequest(c, "Erro ao criar usuário.")
		return
	}
	utils.ResCreated(c)
}

func HandleUpdate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	var requestUser updateUserRequest
	if err := utils.BindJson(c, &requestUser); err != nil {
		return
	}
	if id <= 0 || id != int64(requestUser.Id) {
		utils.ResInvalidParam(c, "id")
		return
	}
	repository := NewRepository()
	user, err := repository.FindById(requestUser.Id)
	if err != nil {
		errorToSearchUser(c)
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
	utils.ResOk(c, "", nil)
}

func HandleDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	if id <= 0 {
		utils.ResInvalidParam(c, "id")
		return
	}

	repository := NewRepository()
	err := repository.Delete(uint(id))
	if err != nil {
		utils.ResBadRequest(c, "Erro ao remover usuário.")
		return
	}
	utils.ResOk(c, "", nil)
}
