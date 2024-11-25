package requests

type UpdateProjectRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
