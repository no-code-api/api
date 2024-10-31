package users

import (
	"strconv"

	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type UserHandler struct {
	DefaultPath string
	userService UserService
}

func NewHandler() UserHandler {
	return UserHandler{
		DefaultPath: "/users",
		userService: NewService(),
	}
}

func (handler UserHandler) HandleLogin(baseHandler *handler.BaseHandler) {
	request := &loginRequest{}
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
	request := &createUserRequest{}
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
	request := &updateUserRequest{}
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
	if err := handler.userService.Delete(uint(id)); err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.Ok("", nil)
}
