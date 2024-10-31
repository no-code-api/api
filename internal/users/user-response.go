package users

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (userResponse *UserResponse) FromModel(user *User) {
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
}
