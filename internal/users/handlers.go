package users

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro-d-santos/no-code-api/internal/jwt"
	"github.com/leandro-d-santos/no-code-api/internal/utils"
)

func errorToSearchUser(c *gin.Context) {
	utils.ResBadRequest(c, "Erro ao consultar usuário.")
}

func HandleLogin(c *gin.Context) {
	userRequest := &loginRequest{}
	if !utils.BindJson(c, userRequest) {
		return
	}

	repository := NewRepository()
	user, ok := repository.FindByEmail(userRequest.Email)
	if !ok {
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

	service := jwt.NewJwtService()
	token, err := service.GenerateJWT(user.Id)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao gerar token.")
		return
	}

	response := &loginResponse{Token: token}
	utils.ResOkData(c, response)
}

func HandleFindAll(c *gin.Context) {
	repository := NewRepository()
	users, ok := repository.FindAll()
	if !ok {
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
	user, ok := repository.FindById(uint(id))
	if !ok {
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
	if !utils.BindJson(c, &user) {
		return
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		utils.ResBadRequest(c, "Erro ao gerar senha.")
	}

	user.Password = hash
	repository := NewRepository()
	if ok := repository.Create(user.ToModel()); !ok {
		utils.ResBadRequest(c, "Erro ao criar usuário.")
		return
	}
	utils.ResCreated(c)
}

func HandleUpdate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	var requestUser updateUserRequest
	if !utils.BindJson(c, &requestUser) {
		return
	}
	if id <= 0 || id != int64(requestUser.Id) {
		utils.ResInvalidParam(c, "id")
		return
	}
	repository := NewRepository()
	user, ok := repository.FindById(requestUser.Id)
	if !ok {
		errorToSearchUser(c)
		return
	}
	if user == nil {
		utils.ResNotFound(c, "Usuário não encontrado.")
		return
	}
	user.Name = requestUser.Name
	if ok := repository.Update(user); !ok {
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

	if _, ok := repository.FindById(uint(id)); !ok {
		utils.ResBadRequest(c, "Usuário não existe.")
		return
	}

	if ok := repository.Delete(uint(id)); !ok {
		utils.ResBadRequest(c, "Erro ao remover usuário.")
		return
	}

	service := jwt.NewJwtService()
	service.RemoveStamp(uint(id))
	utils.ResOk(c, "", nil)
}
