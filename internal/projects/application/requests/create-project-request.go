package requests

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint
}
