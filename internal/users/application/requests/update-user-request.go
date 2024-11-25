package requests

type UpdateUserRequest struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
