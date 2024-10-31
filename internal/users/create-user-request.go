package users

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *createUserRequest) ToModel() *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
