package responses

import "github.com/no-code-api/no-code-api/internal/users/domain/models"

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (userResponse *UserResponse) FromModel(user *models.User) {
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
}
