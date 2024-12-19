package models

type ResourceDynamicFilter struct {
	ProjectId    string
	ResourcePath string
	Fields       []ResourceDynamicFieldFilter
}
