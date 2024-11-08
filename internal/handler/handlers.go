package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type BaseHandler struct {
	context *gin.Context
}

type HandlerFunc func(*BaseHandler)

func NewBaseHandler(context *gin.Context) *BaseHandler {
	return &BaseHandler{
		context: context,
	}
}

func Wrapper(handler HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler(NewBaseHandler(context))
	}
}

func (h *BaseHandler) Param(param string) string {
	return h.context.Param(param)
}

func (h *BaseHandler) Host() string {
	return h.context.Request.Host
}

func (h *BaseHandler) Path() string {
	return h.context.Request.URL.Path
}

func (h *BaseHandler) Method() string {
	return h.context.Request.Method
}

func (h *BaseHandler) Ok(message string, data interface{}) {
	h.Json(http.StatusOK, true, message, data)
}

func (h *BaseHandler) OkData(data interface{}) {
	h.Ok("", data)
}

func (h *BaseHandler) OkMessage(message string) {
	h.Ok(message, nil)
}

func (h *BaseHandler) NotFound(message string) {
	h.Json(http.StatusNotFound, false, message, nil)
}

func (h *BaseHandler) NoContent() {
	h.Json(http.StatusNoContent, true, "", nil)
}

func (h *BaseHandler) Created() {
	h.Json(http.StatusCreated, true, "", nil)
}

func (h *BaseHandler) BadRequest(message string) {
	h.Json(http.StatusBadRequest, false, message, nil)
}

func (h *BaseHandler) InvalidParam(param string) {
	message := fmt.Sprintf("Parâmetro '%v' inválido", param)
	h.Json(http.StatusBadRequest, false, message, nil)
}

func (h *BaseHandler) BindJson(obj any) (success bool) {
	if err := h.context.ShouldBindJSON(obj); err != nil {
		h.BadRequest("Entrada de dados inválida.")
		return false
	}
	return true
}

func (h *BaseHandler) GetUserId() (userId uint, ok bool) {
	userIdContext, exists := h.context.Get("userId")
	if !exists || userIdContext == "" {
		h.BadRequest("Usuário não informado.")
		return 0, false
	}
	return userIdContext.(uint), true
}

func (h *BaseHandler) Json(status int, success bool, message string, data interface{}) {
	h.context.JSON(status, Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	})
}
