package services

import "github.com/leandro-d-santos/no-code-api/internal/external-endpoint/application/requests"

type IExternalEndpointService interface {
	Handle(request *requests.Request) (interface{}, error)
}
