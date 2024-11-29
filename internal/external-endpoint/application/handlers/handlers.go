package handlers

import (
	"fmt"
	"strings"

	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/internal/external-endpoint/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/external-endpoint/application/services"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type ExternalEndpointHandler struct {
	InternalDomain string
	service        services.IExternalEndpointService
}

func NewHandler() ExternalEndpointHandler {
	return ExternalEndpointHandler{
		InternalDomain: config.Env.InternalDomain,
		service:        services.NewService(),
	}
}

// fazer composição de erros
func (handler ExternalEndpointHandler) Handle(baseHandler *handler.BaseHandler) {
	host := baseHandler.Host()
	projectId := handler.GetProjectId(host)
	request := &requests.Request{
		ProjectId: projectId,
		Path:      baseHandler.Path(),
		Method:    baseHandler.Method(),
	}
	data, err := handler.service.Handle(request)
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkData(data)
}

func (handler ExternalEndpointHandler) GetProjectId(host string) string {
	fmt.Println(strings.Split(host, handler.InternalDomain))
	return strings.Split(host, handler.InternalDomain)[0]
}
