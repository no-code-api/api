package projects

type findFilter struct {
	Id     string
	UserId uint
}

type Project struct {
	Id          string
	UserId      uint
	Name        string
	Description string
}
