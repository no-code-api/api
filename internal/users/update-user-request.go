package users

type updateUserRequest struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
