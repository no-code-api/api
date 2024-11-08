package external_endpoint

import (
	"fmt"
	"strings"

	"github.com/leandro-d-santos/no-code-api/config"
	"github.com/leandro-d-santos/no-code-api/internal/handler"
)

type ExternalEndpointHandler struct {
	InternalDomain          string
	externalEndpointService ExternalEndpointService
}

func NewHandler() ExternalEndpointHandler {
	return ExternalEndpointHandler{
		InternalDomain:          config.Env.InternalDomain,
		externalEndpointService: NewService(),
	}
}

// fazer composição de erros
func (handler ExternalEndpointHandler) Handle(baseHandler *handler.BaseHandler) {
	host := baseHandler.Host()
	projectId := handler.GetProjectId(host)
	request := request{
		ProjectId: projectId,
		Path:      baseHandler.Path(),
		Method:    baseHandler.Method(),
	}
	var err error
	method := strings.ToUpper(request.Method)
	switch method {
	case "GET":
		err = handler.externalEndpointService.Get(request)
	case "POST":
		err = handler.externalEndpointService.Post(request)
	default:
		message := fmt.Sprintf("método '%s' não implementado", method)
		baseHandler.BadRequest(message)
		return
	}
	if err != nil {
		baseHandler.BadRequest(err.Error())
		return
	}
	baseHandler.OkMessage("Tamo ai né")
}

func (handler ExternalEndpointHandler) GetProjectId(host string) string {
	fmt.Println(strings.Split(host, handler.InternalDomain))
	return strings.Split(host, handler.InternalDomain)[0]
}
