package utils

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

func ResOk(c *gin.Context, message string, data interface{}) {
	resJson(c, http.StatusOK, true, message, data)
}

func ResOkData(c *gin.Context, data interface{}) {
	ResOk(c, "", data)
}

func ResOkMessage(c *gin.Context, message string) {
	ResOk(c, message, nil)
}

func ResNotFound(c *gin.Context, message string) {
	resJson(c, http.StatusNotFound, false, message, nil)
}

func ResNoContent(c *gin.Context) {
	resJson(c, http.StatusNoContent, true, "", nil)
}

func ResBadRequest(c *gin.Context, message string) {
	resJson(c, http.StatusBadRequest, false, message, nil)
}

func ResInvalidParam(c *gin.Context, param string) {
	message := fmt.Sprintf("Parâmetro '%v' inválido", param)
	resJson(c, http.StatusBadRequest, false, message, nil)
}

func BindJson(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		ResBadRequest(c, "Entrada de dados inválida.")
		return err
	}
	return nil
}

func resJson(c *gin.Context, status int, success bool, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	})
}
