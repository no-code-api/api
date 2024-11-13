package projects

type CreateProjectViewModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint
}
