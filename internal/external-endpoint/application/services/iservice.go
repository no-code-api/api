package services

import "github.com/no-code-api/api/internal/external-endpoint/application/requests"

type IExternalEndpointService interface {
	Handle(request *requests.Request) (interface{}, error)
}
