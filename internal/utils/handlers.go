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
	ResJson(c, http.StatusOK, true, message, data)
}

func ResOkData(c *gin.Context, data interface{}) {
	ResOk(c, "", data)
}

func ResOkMessage(c *gin.Context, message string) {
	ResOk(c, message, nil)
}

func ResNotFound(c *gin.Context, message string) {
	ResJson(c, http.StatusNotFound, false, message, nil)
}

func ResNoContent(c *gin.Context) {
	ResJson(c, http.StatusNoContent, true, "", nil)
}

func ResCreated(c *gin.Context) {
	ResJson(c, http.StatusCreated, true, "", nil)
}

func ResBadRequest(c *gin.Context, message string) {
	ResJson(c, http.StatusBadRequest, false, message, nil)
}

func ResInvalidParam(c *gin.Context, param string) {
	message := fmt.Sprintf("Parâmetro '%v' inválido", param)
	ResJson(c, http.StatusBadRequest, false, message, nil)
}

func BindJson(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		ResBadRequest(c, "Entrada de dados inválida.")
		return err
	}
	return nil
}

func ResJson(c *gin.Context, status int, success bool, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	})
}
