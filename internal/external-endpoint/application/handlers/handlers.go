package handlers

import (
	"fmt"
	"strings"

	"github.com/no-code-api/api/config"
	"github.com/no-code-api/api/internal/external-endpoint/application/requests"
	"github.com/no-code-api/api/internal/external-endpoint/application/services"
	"github.com/no-code-api/api/internal/handler"
	"github.com/no-code-api/api/internal/logger"
)

type ExternalEndpointHandler struct {
	InternalDomain string
	service        services.IExternalEndpointService
	logger         *logger.Logger
}

func NewHandler() ExternalEndpointHandler {
	return ExternalEndpointHandler{
		InternalDomain: config.Env.InternalDomain,
		service:        services.NewService(),
		logger:         logger.NewLogger("ExternalEndpointHandler"),
	}
}

func (handler ExternalEndpointHandler) Handle(baseHandler *handler.BaseHandler) {
	host := baseHandler.Host()
	projectId := handler.GetProjectId(host)
	body, err := handler.getBody(baseHandler)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	request := &requests.Request{
		ProjectId: projectId,
		Path:      baseHandler.Path(),
		Method:    baseHandler.Method(),
		Body:      body,
	}
	data, err := handler.service.Handle(request)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData(data)
}

func (handler ExternalEndpointHandler) GetProjectId(host string) string {
	return strings.Split(host, handler.InternalDomain)[0]
}

func (handler ExternalEndpointHandler) getBody(baseHandler *handler.BaseHandler) (interface{}, error) {
	if baseHandler.Method() != "POST" && baseHandler.Method() != "PUT" {
		return nil, nil
	}
	var body interface{}
	if err := baseHandler.ShouldBindJSON(&body); err != nil {
		return nil, fmt.Errorf("erro ao deserializar o body: %s", err)
	}

	// JSON mimificado - muito mais leve.
	// json, _ := json.Marshal(body)
	// fmt.Println(string(json))
	// fmt.Println(len(json))
	// fmt.Println(float64(len(json)) / 1024)
	return body, nil
}
