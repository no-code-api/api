package users

import (
	"strconv"

	"github.com/leandro-d-santos/no-code-api/internal/handler"
	"github.com/leandro-d-santos/no-code-api/internal/jwt"
)

func errorToSearchUser(h *handler.BaseHandler) {
	h.BadRequest("Erro ao consultar usuário.")
}

func HandleLogin(h *handler.BaseHandler) {
	userRequest := &loginRequest{}
	if !h.BindJson(userRequest) {
		return
	}

	repository := NewRepository()
	user, ok := repository.FindByEmail(userRequest.Email)
	if !ok {
		errorToSearchUser(h)
		return
	}

	if user == nil {
		h.BadRequest("Email inválido.")
		return
	}

	if !VerifyPassword(userRequest.Password, user.Password) {
		h.BadRequest("Senha inválida.")
		return
	}

	service := jwt.NewJwtService()
	token, err := service.GenerateJWT(user.Id)
	if err != nil {
		h.BadRequest("Erro ao gerar token.")
		return
	}

	response := &loginResponse{Token: token}
	h.OkData(response)
}

func HandleFindAll(h *handler.BaseHandler) {
	repository := NewRepository()
	users, ok := repository.FindAll()
	if !ok {
		errorToSearchUser(h)
		return
	}
	usersReponse := make([]*UserResponse, len(users))
	for index, user := range users {
		userResponse := &UserResponse{}
		userResponse.FromModel(user)
		usersReponse[index] = userResponse
	}
	h.OkData(usersReponse)
}

func HandleFindById(h *handler.BaseHandler) {
	id, _ := strconv.ParseInt(h.Param("id"), 10, 32)
	repository := NewRepository()
	user, ok := repository.FindById(uint(id))
	if !ok {
		errorToSearchUser(h)
		return
	}
	if user == nil {
		h.NotFound("Usuário não encontrado.")
		return
	}
	userResponse := &UserResponse{}
	userResponse.FromModel(user)
	h.OkData(userResponse)
}

func HandleCreate(h *handler.BaseHandler) {
	var user createUserRequest
	if !h.BindJson(&user) {
		return
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		h.BadRequest("Erro ao gerar senha.")
	}

	user.Password = hash
	repository := NewRepository()
	if ok := repository.Create(user.ToModel()); !ok {
		h.BadRequest("Erro ao criar usuário.")
		return
	}
	h.Created()
}

func HandleUpdate(h *handler.BaseHandler) {
	id, _ := strconv.ParseInt(h.Param("id"), 10, 32)
	var requestUser updateUserRequest
	if !h.BindJson(&requestUser) {
		return
	}
	if id <= 0 || id != int64(requestUser.Id) {
		h.InvalidParam("id")
		return
	}
	repository := NewRepository()
	user, ok := repository.FindById(requestUser.Id)
	if !ok {
		errorToSearchUser(h)
		return
	}
	if user == nil {
		h.NotFound("Usuário não encontrado.")
		return
	}
	user.Name = requestUser.Name
	if ok := repository.Update(user); !ok {
		h.BadRequest("Erro ao salvar usuário.")
		return
	}
	h.Ok("", nil)
}

func HandleDelete(h *handler.BaseHandler) {
	id, _ := strconv.ParseInt(h.Param("id"), 10, 32)

	if id <= 0 {
		h.InvalidParam("id")
		return
	}
	repository := NewRepository()

	if _, ok := repository.FindById(uint(id)); !ok {
		h.BadRequest("Usuário não existe.")
		return
	}

	if ok := repository.Delete(uint(id)); !ok {
		h.BadRequest("Erro ao remover usuário.")
		return
	}

	service := jwt.NewJwtService()
	service.RemoveStamp(uint(id))
	h.Ok("", nil)
}
