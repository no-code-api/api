package handlers

import (
	"strconv"

	"github.com/no-code-api/no-code-api/internal/handler"
	"github.com/no-code-api/no-code-api/internal/users/application/requests"
	"github.com/no-code-api/no-code-api/internal/users/application/services"
	"github.com/no-code-api/no-code-api/pkg/postgre"
)

type UserHandler struct {
	DefaultPath string
	userService services.IService
}

func NewHandler() UserHandler {
	connection := postgre.GetConnection()
	return UserHandler{
		DefaultPath: "/users",
		userService: services.NewService(connection),
	}
}

func (handler UserHandler) HandleLogin(baseHandler *handler.BaseHandler) {
	request := &requests.LoginRequest{}
	if !baseHandler.BindJson(request) {
		return
	}

	response, err := handler.userService.Login(request)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}

	baseHandler.OkData(response)
}

func (handler UserHandler) HandleFindAll(baseHandler *handler.BaseHandler) {
	users, err := handler.userService.FindAll()
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData(users)
}

func (handler UserHandler) HandleFindById(baseHandler *handler.BaseHandler) {
	id, _ := strconv.ParseInt(baseHandler.Param("id"), 10, 32)
	user, err := handler.userService.FindById(uint(id))
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData(user)
}

func (handler UserHandler) HandleCreate(baseHandler *handler.BaseHandler) {
	request := &requests.CreateUserRequest{}
	if !baseHandler.BindJson(request) {
		return
	}
	if err := handler.userService.Create(request); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Created()
}

func (handler UserHandler) HandleUpdate(baseHandler *handler.BaseHandler) {
	id, _ := strconv.ParseInt(baseHandler.Param("id"), 10, 32)
	request := &requests.UpdateUserRequest{}
	if !baseHandler.BindJson(request) {
		return
	}
	if id <= 0 || id != int64(request.Id) {
		baseHandler.InvalidParam("id")
		return
	}
	if err := handler.userService.Update(request); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Ok("", nil)
}

func (handler UserHandler) HandleDelete(baseHandler *handler.BaseHandler) {
	id, _ := strconv.ParseInt(baseHandler.Param("id"), 10, 32)
	if id <= 0 {
		baseHandler.InvalidParam("id")
		return
	}
	if err := handler.userService.DeleteById(uint(id)); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Ok("", nil)
}
