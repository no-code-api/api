package models

type UpdateResourceDynamic struct {
	ProjectId    string
	ResourcePath string
	Fields       []ResourceDynamicFieldFilter
	Data         interface{}
}
