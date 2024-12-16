package models

type AddResourceDynamic struct {
	ProjectId    string
	ResourcePath string
	Rows         []interface{}
}
