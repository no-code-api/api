package external_endpoint

type ExternalEndpointService struct {
}

func NewService() ExternalEndpointService {
	return ExternalEndpointService{}
}

func (s ExternalEndpointService) Get(request request) error {
	return nil
}

func (s ExternalEndpointService) Post(request request) error {
	return nil
}
