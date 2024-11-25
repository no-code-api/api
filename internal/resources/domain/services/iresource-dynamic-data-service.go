package services

type IResourceDynamicDataService interface {
	CreateCollection(projectId string) error
	DropCollection(projectId string) error
}
